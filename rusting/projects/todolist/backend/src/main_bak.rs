use axum::{
    routing::{get},
    extract::{Path, State},
    http::StatusCode,
    Json, Router,
};
use serde::{Deserialize, Serialize};
use std::sync::{Arc, Mutex};

#[derive(Serialize, Deserialize, Clone)]
struct Todo {
    id: i32,
    title: String,
    completed: bool,
}

type TodoList = Arc<Mutex<Vec<Todo>>>;

#[tokio::main]
async fn main() {
    let todo_list = Arc::new(Mutex::new(Vec::new()));

    // Build the Axum router
    let app = Router::new()
        .route("/", get(|| async { "Hello, world!" }))
        .route("/todos", get(get_all_todos).post(create_todo))
        .route("/todos/:id", get(get_todo))
        .with_state(todo_list);

    // Run the server
    let addr = "127.0.0.1:3000".parse().unwrap();
    println!("Server running at http://{}", addr);
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}

async fn get_all_todos(State(todo_list): State<TodoList>) -> Json<Vec<Todo>> {
    let todos = todo_list.lock().unwrap();
    Json(todos.clone())
}

async fn create_todo(
    State(todo_list): State<TodoList>,
    Json(payload): Json<CreateTodo>,
) -> Json<Todo> {
    let mut todos = todo_list.lock().unwrap();

    let todo = Todo {
        id: todos.len() as i32,
        title: payload.title,
        completed: false,
    };
    todos.push(todo.clone());
    Json(todo)
}

async fn get_todo(
    State(todo_list): State<TodoList>,
    Path(id): Path<i32>,
) -> Result<Json<Todo>, StatusCode> {
    let todos = todo_list.lock().unwrap();

    if let Some(todo) = todos.iter().find(|todo| todo.id == id) {
        Ok(Json(todo.clone()))
    } else {
        Err(StatusCode::NOT_FOUND)
    }
}

#[derive(Deserialize)]
struct CreateTodo {
    title: String,
}
