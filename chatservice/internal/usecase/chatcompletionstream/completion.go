import "github.com/jrsmarcilio/fclx/chatservice/internal/domain/gateway"

type ChatCompletionConfigInputDTO struct {
	Model                string
	ModelMaxTokens       int
	Temperature          float32
	TopP                 float32
	N                    int
	Stop                 []string
	MaxTokens            int
	PresencePenalty      float32
	FrequencyPenalty     float32
	InitialSystemMessage string
}

type ChatCompletionInputDTO struct {
	ChatID      string
	UserID      string
	UserMessage string
	Config      *ChatCompletionConfigInputDTO
}

type ChatCompletionOutputDTO struct {
	ChatID  string
	UserID  string
	Content string
}

type ChatCompletionUseCase struct {
	ChatGateway  gateway.ChatGateway
	OpenAiClient *openai.Client
	Stream 		 chan *ChatCompletionOutputDTO
}

func NewChatCompletionUseCase(chatGateway gateway.ChatGateway, openAiClient *openai.Client, stream chan ChatCompletionOutputDTO) *ChatCompletionUseCase {
	return &ChatCompletionUseCase{
		ChatGateway:  chatGateway,
		OpenAiClient: openAiClient,
	}
}

func (usecase *ChatCompletionUseCase) Execute(context context.Context, input ChatCompletionInputDTO) (*ChatCompletionOutputDTO, error) {
	chat, err := usecase.ChatGateway.FindChatByID(context, input.ChatID)
	if err != nil {
		if err.Error() == "chat not found" {
			// Crie um novo chat (Entity)
			chat, err = createNewChat(input) 
			if err != nil {
				return nil, errors.New("error creating new chat" + err.Error())
			}

			// Salvar chat no banco de dados
			err = usecase.ChatGateway.CreateChat(context, chat)
			if err != nil {
				return nil, errors.New("error persisting new chat" + err.Error())
			}
		} else {
			return nil, errors.New("error fetching existing chat" + err.Error())
		}
	}

	userMessage := entity.NewMessage("user", input.UserMessage, chat.Config.Model)
	if err != nil {
		return nil, errors.New("error creating user message" + err.Error())
	}
	
	// Adicionar mensagem de usuário para o chat (Entity)
	err = chat.AddMessage(userMessage) 
	if err != nil {
		return nil, errors.New("error adding new message to chat" + err.Error())
	}

	// Loop para adicionar todas as mensagens a serem enviadas para o API OpenAI
	messages := []openai.ChatCompletionMessage{}
	for _, msg := range chat.Messages {
		messages = append(messages, openai.ChatCompletionMessage{
			Role: msg.Role,
			Content: msg.Content,
		})
	}

	// Fazendo a chamada para o API OpenAI
	resp, err := usecase.OpenAiClient.CreateChateCompletionStream(
		context,
		openai.ChatCompletionRequest {
			Model: 							chat.Config.Model.Name,
			Messages: 					messages,
			MaxTokens: 					chat.Config.MaxTokens,
			Temperature: 				chat.Config.Temperature,
			TopP: 							chat.Config.TopP,
			PresencePenalty:	 	chat.Config.PresencePenalty,
			FrequencyPenalty: 	chat.Config.FrequencyPenalty,
			Stop: 							chat.Config.Stop,
			Stream: 						true,
		}
	)
	if err != nil {
		return nil, errors.New("error creating chat completion: " + err.Error())
	}

	// Loop para receber a resposta do API OpenAI
	var fullResponse strings.Builder

	for {
		response, err := resp.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, errors.New("error streaming response: " + err.Error())
		}

		// Adicionando a resposta do API OpenAI para o stream
		fullResponse.WriteString(response.Choices[0].Delta.Content)
		r := ChatCompletionOutputDTO{
			ChatID:  chat.ID,
			UserID:  chat.UserID,
			Content: fullResponse.String(),
		}
		// Enviando a resposta para o stream
		usecase.Stream <- &r
	}

	// Criando a mensagem de assistente para o chat (Entity)
	assistant := entity.NewMessage("assistant", fullResponse.String(), chat.Config.Model)
	if err != nil {
		return nil, errors.New("error creating assistant message" + err.Error())
	}
	
	// Adicionar mensagem de assistente para o chat (Entity)
	err = chat.AddMessage(assistant)
	if err != nil {
		return nil, errors.New("error adding new message to chat" + err.Error())
	}

	// Salvar chat no banco de dados
	err = usecase.ChatGateway.SaveChat(context, chat)
	if err != nil {
		return nil, errors.New("error persisting chat" + err.Error())
	}

	// Retornando a resposta para o stream
	return &ChatCompletionOutputDTO{
		ChatID:  chat.ID,
		UserID:  chat.UserID,
		Content: fullResponse.String(),
	}, nil
}



func createNewChat(input ChatCompletionInputDTO) (*entity.Chatm error) {
	model := entity.NewModel(input.Config.Model, input.Config.ModelMaxTokens)

	chatConfig := &entity.ChatConfig {
		FrequencyPenalty: input.Config.FrequencyPenalty,
		MaxTokens:        input.Config.MaxTokens,
		Model:            model,
		N:                input.Config.N,
		PresencePenalty:  input.Config.PresencePenalty,
		Stop:             input.Config.Stop,
		Temperature:      input.Config.Temperature,
		TopP:             input.Config.TopP,
	}


	initialSystemMessage := entity.NewMessage("system", input.Config.InitialSystemMessage, model)
	if err != nil {
		return nil, errors.New("error creating initial system message" + err.Error())
	}

	chat, err := entity.NewChat(input.UserID, initialSystemMessage, chatConfig)
	if err != nil {
		return nil, errors.New("error creating new chat" + err.Error())
	}

	return chat, nil
}