# Charismatic (Team BuzzLightyear)

Charismatic is an Android and Web application designed to assist MSMEsin creating and perfecting promotional materials for their products.
Charismatic enables users to create visually stunning promotional content and evoke emotions with powerful storytelling.
The use of advanced AI technology in Charismatic makes it easier for users to craft narratives that are not only appealing but also effectively conveyed.
As a result, each promotion becomes more persuasive, informative, and captivating to a wide range of consumers.
The technology includes features such as narrative video generator, storyboard maker, natural-toned narrator voice overs, text presentation generator with natural-toned voice overs, and product photo editor.

## Teams

Hustler - **Silvia Larasatul Masyitoh**

Hipster - **Enrico Olivian Maricar**

Hacker - **Widya Ardianto**

## Description

This projects separated in 3 different repositories. Backend, Web App, and Android Mobile App. See the repositories below.

1. [Backend](https://github.com/nartos9090/charismatic-api)
2. [Web App](https://github.com/nartos9090/charismatic-web)
3. [Android Mobile App](https://github.com/Waffle000/Charismatic)

This repository is an API Service written in Go. It uses Gemini AI to generate video storyboard, elevenLabs to generate voice over, and Dall-E to generate storyboard illustration.

Follow this link for the production works. [https://charismatic-api.niwabi.my.id](https://charismatic-api.niwabi.my.id)

## Requirements

1. Mysql Database
2. Go Programming Language. Follow the instructions [here](https://go.dev/doc/install) to install.

## How to use

Create ```.env``` file in the root folder. See the example below.

Run the app

```go run main```

By default, it will run in ```127.0.0.1:8000```.

## API Documentation

Follow the link [https://documenter.getpostman.com/view/9208271/2s9YsRdVAT](https://documenter.getpostman.com/view/9208271/2s9YsRdVAT) to get the API Documentation.

### Testing environment

Create ```.env``` file in the root folder and fill with this configuration.

```
HTTP_PORT=127.0.0.1:8000
JWT_SECRET=qwerty123

DALLE_API_KEY=
ELEVENLABS_API_KEY=
GEMINI_API_KEY=
```

*To prevent github guard check, the API KEY will be provided in the submission file.
