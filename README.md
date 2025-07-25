# Kanona-contract-ai

## GETTING STARTED

First step is to initial lise the backend. Use the following commands to do so:

```
cd backend
go mod init github.com/yourusername/kanona-contract-ai/backend
go get github.com/gorilla/mux
go run main.go
```

You can open your browser and check that the backend is running on the following route:

`http://localhost:8080/health`

Each folder in the backend serves a purpose: 

handlers/	Where your HTTP endpoint logic goes
services/	Where business logic like AI calling or DB interaction lives
models/	Structs defining data shapes (e.g., Contract struct)
utils/	Utility functions like file upload helpers

### Install the MongoDB driver

Firstly install mongodb on your machine. For MacOS use the following commands:

```
brew tap mongodb/brew
brew install mongodb-community
```

Run the following commands in terminal:

```
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
```

Make sure MongoDB is running with the following command in the terminal on windows operating system:
`sudo systemctl start mongod`

Use the following command on macOS: 
`brew services start mongodb/brew/mongodb-community`

And confirm that it is running with the following command:
`brew services list`

You should see the following on your terminal:
```Name                     Status  User     File
mongodb-community@7.0   started yourname ~/Library/LaunchAgents/homebrew.mxcl.mongodb-community@7.0.plist
```
### Installing the Python requirements

For the artificial Intelligent component, python packages will be used.
Once python is installed on your machine, a virtual environment is required and can be initialised with the following commands: 

```
python3 -m venv venv
source venv/bin/activate
```

Also install the necessary packages:

```
pip install flask pymupdf pytesseract spacy
python -m spacy download en_core_web_sm
```
For macOS install tesseract with the following command:

```
brew install tesseract
```
Once you've installed packages in your virtual environment run the following command:
`pip freeze > requirements.txt`

Then you can always reinstall easily via:

`pip install -r requirements.txt`

There is need to setup the python environment in the ai-extraction directory; do so with the following commands in terminal:
```
cd ai-extraction
python3 -m venv venv
source venv/bin/activate
pip install flask pymupdf pytesseract spacy
python -m spacy download en_core_web_sm
pip freeze > requirements.txt
```
