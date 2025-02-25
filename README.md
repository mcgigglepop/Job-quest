# 🏆 Job Application Tracker (JAT)

A full-stack web app that helps you track your job applications, manage interview notes, schedule follow-ups, and automate job searches. Built with **Golang, PostgreSQL, JWT Authentication, and AWS S3**.

## 🚀 Features

✅ **User Authentication** (JWT-based login/signup)  
✅ **Track Job Applications** (Add, edit, delete, filter by status)  
✅ **Notes & Follow-ups** (Set reminders for interviews & follow-ups)  
✅ **Resume Parsing** (Upload resume & extract job-relevant info)  
✅ **Job Scraping API** (Fetch job listings from LinkedIn/Indeed)  
✅ **Email Notifications** (Reminders for follow-ups)  
✅ **RESTful API** (Easily integrate with other apps)  

---

## 📦 Tech Stack

- **Backend**: Golang (Gin/Fiber)
- **Database**: PostgreSQL (GORM)
- **Auth**: JWT-based authentication
- **Storage**: AWS S3 (for resumes)
- **Frontend**: HTMX or React (optional)
- **Deployment**: Docker + AWS/GCP

---

## ⚡ API Endpoints

### 🛡 Authentication
| Method | Endpoint         | Description |
|--------|-----------------|-------------|
| POST   | `/auth/signup`  | Register a new user |
| POST   | `/auth/login`   | Log in and get JWT token |
| GET    | `/auth/me`      | Get current user details |

### 📌 Job Applications
| Method | Endpoint      | Description |
|--------|--------------|-------------|
| POST   | `/jobs`      | Create a new job application |
| GET    | `/jobs`      | Get all jobs (with filters) |
| GET    | `/jobs/{id}` | Get a single job application |
| PUT    | `/jobs/{id}` | Update job details |
| DELETE | `/jobs/{id}` | Delete a job application |

### 📝 Notes & Follow-ups
| Method | Endpoint                | Description |
|--------|-------------------------|-------------|
| POST   | `/jobs/{id}/notes`      | Add a note to a job |
| GET    | `/jobs/{id}/notes`      | Get all notes for a job |
| DELETE | `/jobs/{id}/notes/{id}` | Delete a note |
| POST   | `/jobs/{id}/follow-up`  | Schedule a follow-up reminder |

### 📂 Resume Upload & Parsing
| Method | Endpoint         | Description |
|--------|-----------------|-------------|
| POST   | `/resume/upload` | Upload a resume file (PDF/DOCX) |
| GET    | `/resume/parse`  | Extract job-relevant info from resume |

### 🔍 Job Scraping API
| Method | Endpoint      | Description |
|--------|--------------|-------------|
| GET    | `/jobs/search` | Search jobs from LinkedIn/Indeed |

---

## 🎯 Installation & Setup

1️⃣ Clone the repo:
```sh
git clone https://github.com/yourusername/job-application-tracker.git
cd job-application-tracker
```

2️⃣ Set up environment variables:
```sh
cp .env.example .env
```
Fill in the `.env` file with database credentials, JWT secret, and AWS keys.

3️⃣ Run PostgreSQL (Docker recommended):
```sh
docker-compose up -d
```

4️⃣ Install dependencies:
```sh
go mod tidy
```

5️⃣ Start the server:
```sh
go run main.go
```

6️⃣ (Optional) Run tests:
```sh
go test ./...
```

---

## 📜 License

MIT License. Feel free to fork and modify! 

---

## 🤝 Contributing

PRs are welcome! If you find a bug or have feature suggestions, open an issue.

---

## 📬 Contact

- GitHub: [mcgigglepop](https://github.com/mcgigglepop)
- LinkedIn: [cristopherontiveros](https://linkedin.com/in/cristopherontiveros)


🔥 Happy coding and may you never get ghosted by recruiters again!
