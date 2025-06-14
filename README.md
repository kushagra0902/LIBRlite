# 📚 LIBRlite

A simple simulation of decentralized moderation using **Go**, **PostgreSQL**, and **Goroutines**. This project is designed for learning how to build concurrent backend systems using Go and basic RESTful APIs.

---

## 🚀 What Does LIBRlite Do?

LIBRlite simulates **moderation of messages** using multiple moderators (simulated with goroutines). Each message is processed by 3 moderators who randomly approve or reject it.

You can:

- 📨 Submit a message
- 📤 Fetch messages by timestamp
- 📋 View all messages

Moderation is decided using **majority voting** (2 out of 3).

---

## 🛠 Prerequisites

Before running this project, make sure you have:

✅ Go installed → [https://golang.org/doc/install](https://golang.org/doc/install)  
✅ PostgreSQL installed and running on your local machine  
✅ Git installed (optional but helpful)

---

## ⚙️ Setup Instructions

### 1️⃣ Install PostgreSQL

If you’re on Ubuntu or WSL:

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```

### 2️⃣ Create the Database and Table

```Start the Postgres shell:

bash

sudo -u postgres psql
Then run:

sql

CREATE DATABASE libr;

\c libr

CREATE TABLE messages (
  id UUID PRIMARY KEY,
  content TEXT NOT NULL,
  timestamp BIGINT NOT NULL,
  status VARCHAR(10) NOT NULL
);
Exit with \q.
```

### 3️⃣ Clone the Repository
```
bash
git clone https://github.com/yourusername/LIBRlite.git
cd LIBRlite
```


### 4️⃣ Configure Environment Variables
```
Create a .env file in the root folder:

.env
DB_URL=postgres://your_user:your_password@localhost:5432/libr
Replace your_user and your_password with your local PostgreSQL credentials and 5432(default) with your port number
```

### 5️⃣ Run the Server
```
bash
go mod tidy      # Download dependencies
go build -o server ./cmd
./server
You should see the server running at:
http://localhost:4000
```

### 📡 API Endpoints
```
⚠️ Endpoint names may vary depending on your final implementation.

1)📨 POST /
Send a message for moderation.

Request Body:
{
  "content": "This is a test message"
}

Sample Response:
{
  "id": "generated-uuid",
  "timestamp": 1744219507,
  "status": "approved"
}


2)📥 GET /{timestamp}
Fetch messages by a specific timestamp.


3)📋 GET /
Fetch all messages stored in the system.
```


### 🧠 How Moderation Works
```
Each message is handled by 3 "virtual moderators"

Each moderator is a Go goroutine

They take 1–3 seconds randomly to decide

Their decision is randomly either "approved" or "rejected"

A channel collects the results

If at least 2 say "approved", the message is marked approved

⏱️ A timeout of 3 seconds is applied using context.WithTimeout() to simulate delays.
```


📦 Tech Stack
```
Language	      Go
Database	      PostgreSQL
Router	        gorilla/mux
PG Driver	      pgx
Env Loader	    joho/godotenv
UUIDs	          google/uuid
```

### 🧪 Testing the API
```
Use postman to test the apis

```


This project is for learning/demo purposes only, not for production use.

Only tested on local environments with Postgres running locally.

🧑‍💻 Author
Kushagra Kinra
Feel free to fork the project and customize it!
