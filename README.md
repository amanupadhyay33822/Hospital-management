# Hospital Management System - Backend API

A RESTful API backend for hospital management built with Go and MongoDB, featuring Patient and Doctor modules with appointment scheduling.

## Tech Stack

- **Language**: Go 1.25.5
- **Database**: MongoDB
- **Router**: Gorilla Mux
- **Authentication**: JWT (golang-jwt/jwt)
- **Validation**: go-playground/validator
- **Password Hashing**: bcrypt

## Project Structure

```
test/
├── cmd/server/main.go          # Application entry point
├── internal/
│   ├── auth/                   # JWT token generation
│   ├── config/                 # Database configuration
│   ├── handlers/               # HTTP request handlers
│   ├── middleware/             # Auth & logging middleware
│   ├── models/                 # Data models
│   ├── repository/             # Database operations
│   ├── routes/                 # API route definitions
│   ├── utils/                  # Helper utilities
│   └── validation/             # Input validation schemas
├── uploads/                    # File upload directory
├── go.mod                      # Go module dependencies
└── go.sum                      # Dependency checksums
```

## Features Implemented

### ✅ Authentication
- User registration with email validation
- Login with JWT token generation
- Password hashing with bcrypt
- Auth middleware for protected routes

### ✅ Doctor Management
- Create doctor with duplicate email check
- Get all doctors with pagination
- Get doctor by ID
- Update doctor details
- Delete doctor

### ✅ Appointment System
- Create appointment (requires authentication)
- Get all appointments with filters (doctor_id, patient_id, date)
- Get appointment by ID with doctor details
- Email notifications on booking

### ✅ Utilities
- Async email service
- File upload handling
- Standardized JSON responses
- Request/response logging

## Installation

### Prerequisites
- Go 1.25+ installed
- MongoDB running on localhost:27017
- Git

### Setup

1. **Clone the repository**
```bash
git clone <repository-url>
cd test
```

2. **Install dependencies**
```bash
go mod download
```

3. **Start MongoDB**
```bash
# Make sure MongoDB is running on localhost:27017
mongod
```

4. **Run the application**
```bash
go run cmd/server/main.go
```

Server will start on `http://localhost:8080`

## API Endpoints

### Authentication

#### Register User
```http
POST /register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

#### Login
```http
POST /login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}

Response: { "success": true, "message": "Login successful", "data": "jwt_token" }
```

### Doctor Management

#### Create Doctor
```http
POST /doctors/create
Content-Type: application/json

{
  "name": "Dr. Smith",
  "email": "smith@hospital.com",
  "speciality": "Cardiology"
}
```

#### Get All Doctors
```http
GET /doctors?page=1&limit=10
```

#### Get Doctor by ID
```http
GET /doctors/{id}
```

#### Update Doctor
```http
PUT /doctors/{id}
Content-Type: application/json

{
  "name": "Dr. Smith Updated",
  "speciality": "Neurology"
}
```

#### Delete Doctor
```http
DELETE /doctors/{id}
```

### Appointments

#### Create Appointment (Protected)
```http
POST /appointments
Authorization: Bearer {jwt_token}
Content-Type: application/json

{
  "doctor_id": "507f1f77bcf86cd799439011",
  "appointment_date": "2024-12-25T10:00:00Z"
}
```

#### Get All Appointments
```http
GET /appointments?doctor_id={id}&patient_id={id}&date=2024-12-25
```

#### Get Appointment by ID
```http
GET /appointments/{id}
```

## Database Schema

### User Collection
```json
{
  "_id": "ObjectId",
  "name": "string",
  "email": "string",
  "password": "string (hashed)",
  "role": "string"
}
```

### Doctor Collection
```json
{
  "_id": "ObjectId",
  "name": "string",
  "email": "string",
  "speciality": "string",
  "created_at": "timestamp",
  "updated_at": "timestamp"
}
```

### Appointment Collection
```json
{
  "_id": "ObjectId",
  "patient_id": "ObjectId",
  "doctor_id": "ObjectId",
  "appointment_date": "timestamp",
  "status": "string (booked/cancelled/completed)",
  "created_at": "timestamp",
  "updated_at": "timestamp"
}
```

## Configuration

### Database Connection
Edit `internal/config/database.go`:
```go
Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/test"))
```

### JWT Secret
Edit `internal/auth/jwt.go` to set your JWT secret key.

## Development

### Hot Reload
Using Air for hot reload:
```bash
air
```

Configuration in `.air.toml`

### Build
```bash
go build -o hospital-api cmd/server/main.go
```

### Run Binary
```bash
./hospital-api
```

## Dependencies

```
github.com/gorilla/mux v1.8.1              # HTTP router
github.com/golang-jwt/jwt/v5 v5.3.1        # JWT authentication
go.mongodb.org/mongo-driver v1.17.9        # MongoDB driver
golang.org/x/crypto v0.48.0                # Bcrypt hashing
github.com/go-playground/validator/v10     # Input validation
```

## Pending Features

### Patient Module
- Patient registration and profile management
- Patient medical history
- Patient-appointment relationship

### Enhanced Appointment System
- Update/cancel appointments
- Doctor availability management
- Conflict checking (double booking prevention)
- Time slot management

### Security Enhancements
- Environment variables for sensitive config
- Rate limiting
- CORS configuration
- Password reset functionality
- Email verification
- Refresh tokens

### Advanced Features
- Search and filtering
- Medical records management
- Dashboard analytics
- Prescription management
- Report generation

## Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

## Error Handling

All API responses follow this format:
```json
{
  "success": true/false,
  "message": "Response message",
  "data": {} or null
}
```

## Contributing

1. Fork the repository
2. Create feature branch (`git checkout -b feature/new-feature`)
3. Commit changes (`git commit -am 'Add new feature'`)
4. Push to branch (`git push origin feature/new-feature`)
5. Create Pull Request

## License

MIT License

## Contact

For issues and questions, please open an issue in the repository.

---

**Status**: 🚧 In Development (Core modules 40% complete)
