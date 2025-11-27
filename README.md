# Full Stack Todo List Application

Aplikasi Todo List full-stack yang dibangun dengan React (Frontend) dan Go dengan Fiber framework (Backend), menggunakan PostgreSQL sebagai database.

## 📋 Daftar Isi
- [Fitur yang Diimplementasikan](#-fitur-yang-diimplementasikan)
- [Langkah Instalasi](#-langkah-instalasi)
- [Cara Menjalankan Aplikasi](#-cara-menjalankan-aplikasi)
- [API Documentation](#-api-documentation)
- [Testing](#-menjalankan-tests)
- [Technical Questions](#required-technical-questions)

## 🚀 Fitur yang Diimplementasikan

### Core Features
- ✅ **Todo Management**: Create, Read, Update, Delete todo items
- ✅ **Todo Categories**: Assign dan manage custom categories
- ✅ **Pagination**: Tampil 10 item per halaman dengan kontrol pagination
- ✅ **Search**: Pencarian todo berdasarkan title
- ✅ **Toggle Complete**: Mark todos sebagai completed/incomplete

### Bonus Features
- ✅ **React Context API**: State management menggunakan React Context (+6 points)
- ✅ **Docker**: Containerized backend dan database dengan docker-compose (+3 points)
- ✅ **TypeScript**: Frontend menggunakan TypeScript (+2 points)

## 📦 Langkah Instalasi

### Prerequisites yang Diperlukan
Sebelum memulai, pastikan sistem Anda sudah terinstall:

1. **Docker & Docker Compose**
   ```bash
   # Cek versi Docker
   docker --version
   docker-compose --version
   ```

2. **Node.js (v16 atau lebih tinggi) & npm**
   ```bash
   # Cek versi Node.js
   node --version
   npm --version
   ```

3. **Go (v1.19 atau lebih tinggi)** - Optional, hanya jika ingin run backend tanpa Docker
   ```bash
   # Cek versi Go
   go version
   ```

### Langkah-langkah Instalasi

#### 1. Clone Repository
```bash
git clone https://github.com/alifndaru/allsome-test.git
cd allsome-test
```

#### 2. Setup Environment Variables (Backend)
```bash
cd backend
# Buat file .env jika belum ada (opsional, sudah ada default values)
cp .env.example .env  # Jika ada file example
```

#### 3. Setup Dependencies (Frontend)
```bash
cd frontend
# Install dependencies
npm install

# Atau menggunakan yarn
yarn install
```

## 🚀 Cara Menjalankan Aplikasi

### Opsi 1: Menjalankan dengan Docker (Recommended)

#### Backend + Database (PostgreSQL)
```bash
# Masuk ke folder backend
cd backend

# Jalankan database dan backend API
docker-compose up -d

# Cek status containers
docker-compose ps

# Lihat logs jika ada masalah
docker-compose logs -f
```

Ini akan menjalankan:
- **PostgreSQL database** di port `5432`
- **Go Backend API** di port `3030`
- **Auto migration** dan **seed data** otomatis berjalan

#### Frontend
```bash
# Buka terminal baru, masuk ke folder frontend
cd frontend

# Jalankan development server
npm run dev

# Atau dengan yarn
yarn dev
```

Frontend akan berjalan di `http://localhost:5173`

### Opsi 2: Menjalankan Manual (Tanpa Docker)

#### 1. Setup PostgreSQL Database
```bash
# Install PostgreSQL (macOS dengan Homebrew)
brew install postgresql
brew services start postgresql

# Buat database
createdb todo_db

# Atau bisa menggunakan psql
psql -c "CREATE DATABASE todo_db;"
```

#### 2. Setup Environment Variables
```bash
# Di folder backend, buat/edit file .env
cd backend
cat > .env << EOF
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=todo_db
PORT=3030
EOF
```

#### 3. Jalankan Backend
```bash
cd backend

# Install dependencies
go mod tidy

# Jalankan aplikasi
go run main.go

# Atau build dulu
go build -o main .
./main
```

#### 4. Jalankan Frontend
```bash
cd frontend

# Install dependencies (jika belum)
npm install

# Jalankan development server
npm run dev
```

### Akses Aplikasi

Setelah semua berjalan, Anda bisa akses:

- **Frontend Web App**: http://localhost:5173
- **Backend API**: http://localhost:3030
- **API Base URL**: http://localhost:3030/api/v1
- **Database**: PostgreSQL di localhost:5432 (jika manual setup)

### Verifikasi Instalasi

1. **Cek Backend API**:
   ```bash
   curl http://localhost:3030/api/v1/todos
   ```

2. **Cek Frontend**:
   - Buka browser: http://localhost:5173
   - Seharusnya tampil halaman Todo List

3. **Cek Database**:
   ```bash
   # Jika menggunakan Docker
   docker exec -it todo_postgres psql -U todouser -d todo_db -c "\dt"
   
   # Jika manual setup
   psql -d todo_db -c "\dt"
   ```

## 🧪 Testing

### Backend Testing
Saat ini backend belum memiliki unit tests yang diimplementasikan. Ini adalah salah satu area untuk improvement di masa depan.

**Testing yang bisa dilakukan:**
1. **Manual API Testing** menggunakan cURL atau Postman
2. **Database Connection Testing** - pastikan aplikasi bisa connect ke database
3. **Functional Testing** - test semua endpoints secara manual

### Frontend Testing
Frontend juga belum memiliki automated tests, namun bisa dilakukan testing manual:

1. **Browser Testing** - test UI di berbagai browser
2. **Responsive Testing** - test di berbagai ukuran screen
3. **User Flow Testing** - test complete user journey

### Manual Testing Checklist
- [ ] Backend API berjalan di port 3030
- [ ] Database connection berhasil
- [ ] Frontend berjalan di port 5173
- [ ] CRUD operations untuk todos bekerja
- [ ] CRUD operations untuk categories bekerja
- [ ] Search functionality bekerja
- [ ] Pagination bekerja

## 📚 API Documentation

### Base URL
```
http://localhost:3030/api/v1
```

### Authentication
Saat ini API tidak memerlukan authentication. Semua endpoints bisa diakses secara langsung.

---

## 📝 Todos API

### 1. List Todos
**Endpoint:** `GET /todos`

**Description:** Mengambil daftar todos dengan pagination dan search

**Query Parameters:**
| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `page` | integer | 1 | Nomor halaman |
| `limit` | integer | 10 | Jumlah item per halaman |
| `search` | string | "" | Pencarian berdasarkan title (case-insensitive) |

**Example Request:**
```bash
curl "http://localhost:3030/api/v1/todos?page=1&limit=10&search=coding"
```

**Example Response:**
```json
{
  "status": "success",
  "message": "todos found",
  "data": {
    "items": [
      {
        "id": 1,
        "title": "Complete coding challenge",
        "description": "Build a full-stack todo application for Industrix",
        "completed": false,
        "category_id": 2,
        "category": {
          "id": 2,
          "name": "Work", 
          "color": "#3B82F6"
        },
        "priority": "high",
        "due_date": "2024-08-03T23:59:59Z",
        "created_at": "2024-07-31T10:00:00Z",
        "updated_at": "2024-07-31T10:00:00Z"
      }
    ],
    "pagination": {
      "current_page": 1,
      "per_page": 10,
      "total": 25,
      "total_pages": 3
    }
  }
}
```

### 2. Get Todo by ID
**Endpoint:** `GET /todos/:id`

**Description:** Mengambil detail todo berdasarkan ID

**Example Request:**
```bash
curl http://localhost:3030/api/v1/todos/1
```

**Example Response:**
```json
{
  "status": "success",
  "message": "todo found",
  "data": {
    "id": 1,
    "title": "Complete coding challenge",
    "description": "Build a full-stack todo application",
    "completed": false,
    "category_id": 2,
    "category": {
      "id": 2,
      "name": "Work",
      "color": "#3B82F6"
    },
    "priority": "high",
    "due_date": "2024-08-03T23:59:59Z",
    "created_at": "2024-07-31T10:00:00Z",
    "updated_at": "2024-07-31T10:00:00Z"
  }
}
```

### 3. Create Todo
**Endpoint:** `POST /todos`

**Description:** Membuat todo baru

**Request Body:**
```json
{
  "title": "Complete coding challenge",
  "description": "Build a full-stack todo application",
  "category_id": 1,
  "priority": "high",
  "due_date": "2024-08-03T23:59:59Z"
}
```

**Required Fields:**
- `title` (string): Judul todo
- `description` (string): Deskripsi todo
- `category_id` (integer): ID kategori
- `priority` (string): Level prioritas ("high", "medium", "low")

**Optional Fields:**
- `due_date` (string): Tanggal deadline dalam format ISO 8601

**Example Request:**
```bash
curl -X POST http://localhost:3030/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{
    "title": "New Todo",
    "description": "Todo description",
    "category_id": 1,
    "priority": "medium"
  }'
```

**Example Response:**
```json
{
  "status": "success",
  "message": "todo created",
  "data": {
    "id": 5,
    "title": "New Todo",
    "description": "Todo description", 
    "completed": false,
    "category_id": 1,
    "priority": "medium",
    "due_date": null,
    "created_at": "2024-11-27T10:00:00Z",
    "updated_at": "2024-11-27T10:00:00Z"
  }
}
```

### 4. Update Todo
**Endpoint:** `PUT /todos/:id`

**Description:** Update todo yang sudah ada

**Request Body:** (sama seperti create todo)
```json
{
  "title": "Updated Todo Title",
  "description": "Updated description",
  "category_id": 2,
  "priority": "low",
  "completed": true,
  "due_date": "2024-08-05T23:59:59Z"
}
```

**Example Request:**
```bash
curl -X PUT http://localhost:3030/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Todo",
    "description": "Updated description",
    "category_id": 1,
    "priority": "low",
    "completed": true
  }'
```

### 5. Delete Todo
**Endpoint:** `DELETE /todos/:id`

**Description:** Menghapus todo berdasarkan ID

**Example Request:**
```bash
curl -X DELETE http://localhost:3030/api/v1/todos/1
```

**Example Response:**
```json
{
  "status": "success",
  "message": "todo deleted"
}
```

### 6. Toggle Todo Completion
**Endpoint:** `PATCH /todos/:id/complete`

**Description:** Toggle status completion todo

**Request Body:**
```json
{
  "completed": true
}
```

**Example Request:**
```bash
curl -X PATCH http://localhost:3030/api/v1/todos/1/complete \
  -H "Content-Type: application/json" \
  -d '{"completed": true}'
```

---

## 🏷 Categories API

### 1. List Categories
**Endpoint:** `GET /categories`

**Description:** Mengambil semua kategori

**Example Request:**
```bash
curl http://localhost:3030/api/v1/categories
```

**Example Response:**
```json
{
  "status": "success", 
  "message": "categories found",
  "data": [
    {
      "id": 1,
      "name": "Work",
      "color": "#3B82F6",
      "created_at": "2024-07-31T10:00:00Z",
      "updated_at": "2024-07-31T10:00:00Z"
    },
    {
      "id": 2,
      "name": "Personal",
      "color": "#F97316",
      "created_at": "2024-07-31T10:00:00Z", 
      "updated_at": "2024-07-31T10:00:00Z"
    }
  ]
}
```

### 2. Create Category
**Endpoint:** `POST /categories`

**Description:** Membuat kategori baru

**Request Body:**
```json
{
  "name": "Shopping",
  "color": "#10B981"
}
```

**Required Fields:**
- `name` (string): Nama kategori (unique)

**Optional Fields:**
- `color` (string): Warna dalam format hex

**Example Request:**
```bash
curl -X POST http://localhost:3030/api/v1/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Shopping",
    "color": "#10B981"
  }'
```

### 3. Update Category
**Endpoint:** `PUT /categories/:id`

**Description:** Update kategori yang sudah ada

**Request Body:** (sama seperti create category)

### 4. Delete Category
**Endpoint:** `DELETE /categories/:id`

**Description:** Menghapus kategori berdasarkan ID

---

## 🚨 Error Responses

Semua error response menggunakan format standar:

```json
{
  "status": "error",
  "message": "Error message",
  "details": "Detailed error information"
}
```

**HTTP Status Codes:**
- `200` - Success
- `201` - Created
- `400` - Bad Request (validation error)
- `404` - Not Found
- `500` - Internal Server Error

**Common Error Examples:**

**Validation Error (400):**
```json
{
  "status": "error",
  "message": "validation failed",
  "details": "title is required"
}
```

**Not Found (404):**
```json
{
  "status": "error",
  "message": "todo not found",
  "details": "record not found"
}
```

---

## 🧪 Testing API

### Menggunakan cURL
Semua contoh di atas menggunakan cURL untuk testing API.

### Menggunakan Postman
1. Import collection dari file `postman_collection.json` (jika tersedia)
2. Set base URL: `http://localhost:3030/api/v1`
3. Test semua endpoints sesuai dokumentasi

### Menggunakan VS Code REST Client
Buat file `test.http`:
```http
### Get all todos
GET http://localhost:3030/api/v1/todos

### Create new todo  
POST http://localhost:3030/api/v1/todos
Content-Type: application/json

{
  "title": "Test Todo",
  "description": "Test description", 
  "category_id": 1,
  "priority": "medium"
}

### Get categories
GET http://localhost:3030/api/v1/categories
```

---

# Required Technical Questions

## Database Design Questions

### 1. What database tables did you create and why?

#### Tabel yang dibuat:

**a. `categories` table**
```sql
CREATE TABLE categories (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    color VARCHAR(20),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);
```
- **Purpose**: Menyimpan kategori-kategori todo (Work, Personal, Shopping, dll)
- **Fields**: 
  - `id`: Primary key auto-increment
  - `name`: Nama kategori (unique untuk mencegah duplikasi)
  - `color`: Warna untuk display di UI
  - `created_at`, `updated_at`: Timestamps untuk audit

**b. `todos` table**
```sql
CREATE TABLE todos (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    category_id BIGINT,
    priority VARCHAR(10),
    due_date TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);
```
- **Purpose**: Menyimpan data todo items utama
- **Fields**:
  - `id`: Primary key auto-increment
  - `title`: Judul todo (required)
  - `description`: Deskripsi detail todo
  - `completed`: Status completion (default false)
  - `category_id`: Foreign key ke table categories
  - `priority`: Level prioritas (high, medium, low)
  - `due_date`: Tanggal deadline
  - `created_at`, `updated_at`: Timestamps untuk audit

#### Relasi antar tabel:
- **One-to-Many**: Satu category bisa memiliki banyak todos
- **Foreign Key**: `todos.category_id` references `categories.id`
- **Optional Relationship**: Todo bisa tidak memiliki kategori (category_id nullable)

#### Mengapa struktur ini dipilih:
1. **Normalisasi**: Memisahkan categories dan todos menghindari redundansi data
2. **Flexibility**: Category terpisah memungkinkan easy management dan reusability
3. **Scalability**: Struktur ini mudah di-extend untuk fitur future (tags, users, etc.)
4. **Performance**: Simple relationship yang efficient untuk queries

### 2. How did you handle pagination and filtering in the database?

#### Pagination Implementation:
```go
func (r *todoRepository) List(offset int, limit int, search string) ([]models.Todos, int64, error) {
    var todos []models.Todos
    var total int64

    query := r.db.Model(&models.Todos{}).Preload("Category")
    
    // Apply search filter if provided
    if search != "" {
        like := "%" + search + "%"
        query = query.Where("title ILIKE ?", like)
    }
    
    // Count total records for pagination
    if err := query.Count(&total).Error; err != nil {
        return nil, 0, err
    }
    
    // Apply pagination and ordering
    if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&todos).Error; err != nil {
        return nil, 0, err
    }
    
    return todos, total, nil
}
```

#### Queries untuk filtering dan sorting:
1. **Search Query**: Menggunakan `ILIKE` untuk case-insensitive search pada field `title`
2. **Pagination Query**: Menggunakan `OFFSET` dan `LIMIT` untuk pagination
3. **Sorting**: Default `ORDER BY created_at DESC` untuk menampilkan todo terbaru first
4. **Preloading**: `Preload("Category")` untuk eager loading relasi category

#### Efficient Pagination Handling:
- **Count Query**: Terpisah untuk mendapatkan total records tanpa load semua data
- **Offset-Limit**: Standard pagination approach yang efficient untuk small-medium datasets
- **Single Query**: Menggunakan GORM untuk minimize database calls

#### Indexes yang ditambahkan:
Meskipun tidak explicitly dibuat custom index, PostgreSQL secara otomatis membuat:
- **Primary Key Index**: Pada `id` columns untuk fast lookups
- **Foreign Key Index**: Pada `category_id` untuk efficient joins

**Future optimization yang bisa ditambahkan:**
```sql
CREATE INDEX idx_todos_title ON todos USING gin(to_tsvector('english', title));
CREATE INDEX idx_todos_created_at ON todos(created_at DESC);
CREATE INDEX idx_todos_category_id ON todos(category_id);
```

## Technical Decision Questions

### 1. How did you implement responsive design?

#### Breakpoints yang digunakan:
Menggunakan **Ant Design's built-in responsive system** yang memiliki breakpoints:
- `xs`: < 576px (Mobile)
- `sm`: ≥ 576px (Small devices)
- `md`: ≥ 768px (Tablets) 
- `lg`: ≥ 992px (Desktop)
- `xl`: ≥ 1200px (Large desktop)

#### UI Adaptation di berbagai screen sizes:

**Desktop (lg+):**
```tsx
// Layout dengan header horizontal menu
<Layout>
  <Header>
    <Menu theme="dark" mode="horizontal" items={menuItems} />
  </Header>
  <Content style={{ padding: "20px" }}>
    {children}
  </Content>
</Layout>
```

**Tablet/Mobile adaptations:**
- **Table Component**: Ant Design Table otomatis menjadi scrollable horizontally di mobile
- **Form Layout**: Modal forms tetap readable di mobile dengan vertical layout
- **Button Spacing**: Menggunakan Ant Design's spacing system yang responsive
- **Input Fields**: Full width di mobile devices

#### Ant Design components yang membantu responsiveness:

1. **Layout Components**:
   ```tsx
   <Layout> // Responsive layout container
   <Header> // Auto-adapts menu orientation
   <Content> // Flexible content area
   ```

2. **Grid System**:
   ```tsx
   <Row gutter={[16, 16]}>
     <Col xs={24} sm={12} md={8} lg={6}>
       // Responsive columns
     </Col>
   </Row>
   ```

3. **Table Component**:
   ```tsx
   <Table 
     scroll={{ x: 'max-content' }} // Horizontal scroll on mobile
     size="middle" // Compact size for mobile
   />
   ```

4. **Form Components**:
   ```tsx
   <Form layout="vertical"> // Vertical labels for mobile
     <Form.Item>
       <Input /> // Auto-sizing inputs
     </Form.Item>
   </Form>
   ```

#### CSS Adaptations:
```css
/* Menggunakan Ant Design variables untuk consistency */
.ant-layout-content {
  padding: 20px; /* Desktop */
}

@media (max-width: 768px) {
  .ant-layout-content {
    padding: 12px; /* Mobile */
  }
}
```

### 2. How did you structure your React components?

#### Component Hierarchy:
```
App.tsx (Root)
├── BrowserRouter
├── TodoProvider (Context)
├── CategoryProvider (Context)
└── AppLayout
    ├── Header (Navigation)
    └── Content
        └── Routes
            ├── TodoListPage
            │   ├── TodoTable
            │   └── TodoForm
            └── CategoryPage
```

#### State Management between components:

**1. React Context API Implementation:**
```tsx
// TodoContext.tsx
export const TodoContext = createContext<TodoContextType>({
  todos: [],
  setTodos: () => {},
});

export const TodoProvider = ({ children }: { children: ReactNode }) => {
  const [todos, setTodos] = useState<Todo[]>([]);
  return (
    <TodoContext.Provider value={{ todos, setTodos }}>
      {children}
    </TodoContext.Provider>
  );
};
```

**2. Custom Hooks untuk State Management:**
```tsx
// useTodos.ts
export const useTodos = (page: number, search: string) => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    // Fetch todos from API
    fetchTodos();
  }, [page, search]);

  return { todos, loading };
};
```

**3. Component Communication:**
- **Parent to Child**: Props passing
- **Child to Parent**: Callback functions
- **Sibling Components**: Shared context state
- **Global State**: React Context API

#### Filtering dan Pagination State Handling:

**TodoListPage.tsx:**
```tsx
export default function TodoListPage() {
  const [page] = useState(1);
  const [search, setSearch] = useState("");
  
  // Custom hook handles API calls and state
  const { todos } = useTodos(page, search);
  
  return (
    <>
      <Input.Search
        value={search}
        onChange={(e) => setSearch(e.target.value)} // Local state update
      />
      <TodoTable data={todos} /> // Pass data via props
    </>
  );
}
```

**Benefits dari struktur ini:**
- **Separation of Concerns**: Each component has single responsibility
- **Reusability**: Components bisa digunakan di multiple places
- **Maintainability**: Easy to modify individual components
- **Type Safety**: Full TypeScript support

### 3. What backend architecture did you choose and why?

#### Architecture Pattern: **Clean Architecture dengan Layer Separation**

```
Controllers (HTTP Layer)
    ↓
Services (Business Logic Layer)
    ↓
Repositories (Data Access Layer)
    ↓
Models (Database Layer)
```

#### API Routes Organization:
```go
// routes/routes.go
func Setup(app *fiber.App) {
    api := app.Group("/api/v1")
    
    // Todo Routes
    todos := api.Group("/todos")
    todos.Get("/", todoController.List)           // GET /api/v1/todos
    todos.Post("/", todoController.Create)        // POST /api/v1/todos
    todos.Get("/:id", todoController.GetByID)     // GET /api/v1/todos/:id
    todos.Put("/:id", todoController.Update)      // PUT /api/v1/todos/:id
    todos.Delete("/:id", todoController.Delete)   // DELETE /api/v1/todos/:id
    todos.Patch("/:id/complete", todoController.ToggleComplete)
    
    // Categories Routes
    categories := api.Group("/categories")
    categories.Get("/", categoryController.ListCategories)
    categories.Post("/", categoryController.CreateCategory)
    // ... etc
}
```

#### Code Structure Organization:

**1. Controllers Layer** - HTTP Request Handling:
```go
type TodoController struct {
    services services.TodoService
}

func (ctl *TodoController) Create(c *fiber.Ctx) error {
    // Parse request
    var todoInput request.CreateTodoRequest
    if err := c.BodyParser(&todoInput); err != nil {
        return utils.BadRequest(c, "Gagal Parsing Data", err.Error())
    }
    
    // Call service layer
    err := ctl.services.Create(&todo)
    if err != nil {
        return utils.BadRequest(c, "create todo failed", err.Error())
    }
    return utils.Created(c, "todo created", todo)
}
```

**2. Services Layer** - Business Logic:
```go
type todoService struct {
    repo repositories.TodoRepository
}

func (s *todoService) Create(input *models.Todos) error {
    // Business logic validation
    if input.Title == "" {
        return errors.New("title is required")
    }
    if input.CategoryID == 0 {
        return errors.New("category_id is required")
    }
    
    // Call repository layer
    return s.repo.Create(input)
}
```

**3. Repositories Layer** - Data Access:
```go
type todoRepository struct {
    db *gorm.DB
}

func (r *todoRepository) Create(todo *models.Todos) error {
    return r.db.Create(todo).Error
}

func (r *todoRepository) List(offset int, limit int, search string) ([]models.Todos, int64, error) {
    // Database queries with GORM
    // ... implementation
}
```

#### Error Handling Approach:

**1. Centralized Error Responses:**
```go
// utils/response.go
func BadRequest(c *fiber.Ctx, message string, details string) error {
    return c.Status(400).JSON(fiber.Map{
        "status":  "error",
        "message": message,
        "details": details,
    })
}

func Success(c *fiber.Ctx, message string, data interface{}) error {
    return c.Status(200).JSON(fiber.Map{
        "status":  "success",
        "message": message,
        "data":    data,
    })
}
```

**2. Layer-by-Layer Error Propagation:**
- **Repository Layer**: Return raw database errors
- **Service Layer**: Add business context to errors
- **Controller Layer**: Convert to HTTP responses

**3. Validation Errors:**
```go
// Service layer validation
if input.Title == "" {
    return errors.New("title is required")
}

// Controller layer response
if err != nil {
    return utils.BadRequest(c, "validation failed", err.Error())
}
```

#### Why this architecture?

**Advantages:**
1. **Separation of Concerns**: Each layer has specific responsibility
2. **Testability**: Easy to unit test each layer independently
3. **Maintainability**: Easy to modify business logic without touching HTTP layer
4. **Scalability**: Easy to add new features or modify existing ones
5. **Dependency Injection**: Loose coupling between layers

### 4. How did you handle data validation?

#### Validation Strategy: **Multi-Layer Validation**

**Where validation occurs:**
- ✅ **Frontend Validation**: UI/UX improvement dan immediate feedback
- ✅ **Backend Validation**: Security dan data integrity
- ✅ **Database Constraints**: Final safety net

#### Frontend Validation (React + Ant Design):

**Form-level validation:**
```tsx
<Form form={form} onFinish={onSubmit}>
  <Form.Item 
    name="title" 
    label="Title" 
    rules={[
      { required: true, message: 'Title is required!' },
      { min: 3, message: 'Title must be at least 3 characters!' },
      { max: 255, message: 'Title cannot exceed 255 characters!' }
    ]}
  >
    <Input />
  </Form.Item>
  
  <Form.Item 
    name="category_id" 
    label="Category"
    rules={[{ required: true, message: 'Please select a category!' }]}
  >
    <Select>
      {categories.map(c => (
        <Select.Option key={c.id} value={c.id}>{c.name}</Select.Option>
      ))}
    </Select>
  </Form.Item>
</Form>
```

**Real-time validation:**
- **Required fields**: Immediate visual feedback
- **Format validation**: Character limits, valid formats
- **Business rules**: Category selection required

#### Backend Validation (Go Services Layer):

**Business Logic Validation:**
```go
func (s *todoService) Create(input *models.Todos) error {
    // Required field validation
    if input.Title == "" {
        return errors.New("title is required")
    }
    
    if input.Description == "" {
        return errors.New("description is required")
    }
    
    if input.CategoryID == 0 {
        return errors.New("category_id is required")
    }
    
    if input.Priority == "" {
        return errors.New("priority is required")
    }
    
    // Business rule validation
    if input.Priority != "high" && input.Priority != "medium" && input.Priority != "low" {
        return errors.New("priority must be high, medium, or low")
    }
    
    return s.repo.Create(input)
}
```

**Request Struct Validation:**
```go
type CreateTodoRequest struct {
    Title       string     `json:"title" validate:"required,min=1,max=255"`
    Description string     `json:"description" validate:"max=1000"`
    CategoryID  int64      `json:"category_id" validate:"required,min=1"`
    Priority    string     `json:"priority" validate:"required,oneof=high medium low"`
    DueDate     *time.Time `json:"due_date"`
}
```

#### Database-Level Constraints:

**Schema Constraints:**
```sql
CREATE TABLE todos (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,        -- NOT NULL constraint
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    category_id BIGINT,                 -- Foreign key constraint
    priority VARCHAR(10),
    due_date TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    
    -- Foreign key constraint
    CONSTRAINT fk_todos_category 
        FOREIGN KEY (category_id) 
        REFERENCES categories(id)
);

CREATE TABLE categories (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,  -- UNIQUE constraint
    color VARCHAR(20),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);
```

#### Validation Rules Implementation:

**Todo Validation Rules:**
1. **Title**: Required, 1-255 characters
2. **Description**: Optional, max 1000 characters  
3. **Category**: Required, must exist in categories table
4. **Priority**: Required, must be 'high', 'medium', or 'low'
5. **Due Date**: Optional, must be valid timestamp
6. **Completed**: Boolean, defaults to false

**Category Validation Rules:**
1. **Name**: Required, unique, 1-100 characters
2. **Color**: Optional, valid color format

#### Why Multi-Layer Validation?

**Frontend Benefits:**
- **User Experience**: Immediate feedback, no server round-trip
- **Reduce Server Load**: Catch basic errors before API calls
- **Visual Guidance**: Clear error messages and field highlighting

**Backend Benefits:**
- **Security**: Don't trust client-side validation
- **Data Integrity**: Ensure business rules are enforced
- **API Consistency**: Same validation for all API consumers

**Database Benefits:**
- **Final Safety Net**: Prevent invalid data at storage level
- **Referential Integrity**: Foreign key constraints
- **Performance**: Database constraints are highly optimized

## Testing & Quality Questions

### 1. What did you choose to unit test and why?

#### Current Testing Status:

**Saat ini belum ada unit tests yang diimplementasikan** dalam project ini karena keterbatasan waktu dalam development. Namun berikut adalah strategy testing yang akan digunakan jika ada waktu lebih:

**Backend Testing Strategy:**

**1. Repository Layer Tests:**
```go
// todo_repository_test.go
func TestTodoRepository_Create(t *testing.T) {
    // Test cases:
    // - Valid todo creation
    // - Invalid data handling
    // - Database constraint violations
}

func TestTodoRepository_List(t *testing.T) {
    // Test cases:
    // - Pagination functionality
    // - Search filtering
    // - Empty result sets
    // - Large datasets performance
}
```

**2. Service Layer Tests:**
```go
// todo_service_test.go  
func TestTodoService_Create(t *testing.T) {
    // Test cases:
    // - Valid input validation
    // - Required field validation
    // - Business rule enforcement
    // - Error propagation from repository
}

func TestTodoService_ListTodos(t *testing.T) {
    // Test cases:
    // - Pagination calculations
    // - Search term processing
    // - Empty search results
}
```

**3. Controller Layer Tests:**
```go
// todo_controller_test.go
func TestTodoController_Create(t *testing.T) {
    // Test cases:
    // - Valid HTTP requests
    // - Malformed JSON handling
    // - Authentication/authorization
    // - HTTP status codes
    // - Response format consistency
}
```

#### Edge Cases yang dipertimbangkan:

**1. Data Validation Edge Cases:**
- Empty/null values untuk required fields
- Extremely long strings (beyond limits)
- Invalid data types
- SQL injection attempts
- XSS attempts in input fields

**2. Pagination Edge Cases:**
- Page number = 0 atau negative
- Limit exceed maximum allowed
- Empty result sets
- Single item result sets

**3. Search Edge Cases:**
- Empty search queries
- Special characters dalam search
- Very long search terms
- Case sensitivity testing

**4. Database Edge Cases:**
- Database connection failures
- Transaction rollbacks
- Concurrent access scenarios
- Foreign key constraint violations

#### Test Structure:

**Testing Framework Setup:**
```go
// setup_test.go
func setupTestDB() *gorm.DB {
    // In-memory SQLite untuk testing
    db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    db.AutoMigrate(&models.Todo{}, &models.Category{})
    return db
}

func TestMain(m *testing.M) {
    // Setup test environment
    // Run tests
    // Cleanup
}
```

**Mock Dependencies:**
```go
type mockTodoRepository struct{}

func (m *mockTodoRepository) Create(todo *models.Todos) error {
    // Mock implementation untuk service layer testing
}

func TestTodoService_WithMocks(t *testing.T) {
    service := NewTodoService(&mockTodoRepository{})
    // Test business logic without database dependencies
}
```

#### Frontend Testing Strategy:

**Component Testing:**
```tsx
// TodoTable.test.tsx
describe('TodoTable Component', () => {
  test('renders todo items correctly', () => {
    // Test data rendering
  });
  
  test('handles empty data state', () => {
    // Test empty state display
  });
  
  test('triggers edit callback on edit button click', () => {
    // Test user interactions
  });
});
```

**Hook Testing:**
```tsx
// useTodos.test.tsx
describe('useTodos Hook', () => {
  test('fetches todos on mount', () => {
    // Test API calls
  });
  
  test('updates todos on search change', () => {
    // Test search functionality
  });
});
```

### 2. If you had more time, what would you improve or add?

#### Technical Debt yang akan diatasi:

**1. Comprehensive Testing (Priority #1):**
- **Backend Unit Tests**: Implementasi testing untuk service dan repository layer (+10 bonus points)
  - Repository layer tests untuk database operations
  - Service layer tests untuk business logic
  - Controller tests untuk HTTP handling
- **Frontend Testing**: Component, hook, dan integration tests dengan Jest/React Testing Library
- **E2E Testing**: Cypress atau Playwright untuk full user journey testing
- **API Testing**: Postman collections atau automated API tests dengan Newman

**2. Advanced Error Handling:**
```go
// Custom error types untuk better error handling
type ValidationError struct {
    Field   string `json:"field"`
    Message string `json:"message"`
}

type APIError struct {
    Code    string             `json:"code"`
    Message string             `json:"message"`
    Details []ValidationError  `json:"details,omitempty"`
}
```

**3. Database Optimizations:**
- **Database Indexing**: Proper indexes untuk search dan pagination
- **Query Optimization**: Analyze dan optimize slow queries
- **Connection Pooling**: Better database connection management
- **Migrations**: More robust migration system

**4. Security Improvements:**
- **Input Sanitization**: Prevent XSS dan SQL injection
- **Rate Limiting**: Prevent API abuse
- **CORS Configuration**: Proper CORS setup
- **Environment Security**: Secrets management

#### Features yang akan ditambahkan:

**1. Advanced Filtering (+5 bonus points):**
```tsx
// Advanced filter component
const FilterPanel = () => (
  <Space direction="vertical">
    <Select placeholder="Filter by Status">
      <Option value="completed">Completed</Option>
      <Option value="pending">Pending</Option>
    </Select>
    
    <Select placeholder="Filter by Category">
      {categories.map(c => (
        <Option key={c.id} value={c.id}>{c.name}</Option>
      ))}
    </Select>
    
    <Select placeholder="Filter by Priority">
      <Option value="high">High</Option>
      <Option value="medium">Medium</Option>
      <Option value="low">Low</Option>
    </Select>
  </Space>
);
```

**2. User Management & Authentication:**
- User registration/login system
- Todo ownership (user-specific todos)
- Role-based access control
- Social login integration

**3. Real-time Features:**
- WebSocket integration untuk real-time updates
- Collaborative editing
- Live notifications
- Activity feeds

**4. Performance Enhancements:**
- **Frontend**: Virtual scrolling untuk large lists
- **Backend**: Caching layer (Redis)
- **Database**: Query optimization dan indexing
- **CDN**: Static asset optimization

**5. Advanced UI/UX:**
- **Drag & Drop**: Reorder todos with drag and drop
- **Bulk Operations**: Select multiple todos untuk bulk actions
- **Keyboard Shortcuts**: Power user features
- **Dark Mode**: Theme switching
- **Accessibility**: Full WCAG compliance

**6. Mobile Experience:**
- **PWA**: Progressive Web App capabilities
- **Offline Support**: Work offline with sync when online
- **Mobile-First Design**: Better mobile UX
- **Push Notifications**: Mobile notifications

#### Architecture Improvements:

**1. Microservices Architecture:**
```
API Gateway
├── User Service
├── Todo Service  
├── Category Service
└── Notification Service
```

**2. Event-Driven Architecture:**
- Event sourcing untuk audit trails
- Message queues untuk asynchronous processing
- CQRS untuk read/write separation

**3. Monitoring & Observability:**
- **Logging**: Structured logging dengan correlation IDs
- **Metrics**: Application performance metrics
- **Tracing**: Distributed tracing
- **Health Checks**: Service health monitoring

**4. DevOps Improvements:**
- **CI/CD Pipeline**: Automated testing dan deployment
- **Environment Management**: Staging, production environments
- **Infrastructure as Code**: Terraform atau similar
- **Container Orchestration**: Kubernetes deployment

#### What would be refactored:

**1. Code Organization:**
- **Folder Structure**: More modular organization
- **Code Splitting**: Lazy loading untuk better performance
- **Shared Components**: Component library
- **Type Definitions**: More comprehensive TypeScript types

**2. State Management:**
- **Redux Toolkit**: More sophisticated state management
- **React Query**: Better server state management
- **Optimistic Updates**: Better user experience

**3. API Design:**
- **GraphQL**: More flexible API queries  
- **API Versioning**: Proper versioning strategy
- **OpenAPI Spec**: Complete API documentation
- **Batch Operations**: Reduce API calls

---