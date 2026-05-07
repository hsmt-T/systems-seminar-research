use axum::{routing::get, Json, Router};
use db::{create_pool, db_health, DbPool};
use serde_json::{json, Value};
use std::{env, net::SocketAddr};

mod db;

#[derive(Clone)]
struct AppState {
    db_pool: DbPool,
}

#[tokio::main]
async fn main() {
    let database_url = env::var("DATABASE_URL")
        .unwrap_or_else(|_| "postgres://postgres:postgres@localhost:5432/app".to_string());
    let db_pool = create_pool(&database_url)
        .await
        .expect("failed to create database pool");
    let state = AppState { db_pool };

    let app = Router::new()
        .route("/", get(root))
        .route("/health", get(health))
        .route("/db-health", get(db_health))
        .with_state(state);

    let addr = SocketAddr::from(([0, 0, 0, 0], 3000));
    let listener = tokio::net::TcpListener::bind(addr)
        .await
        .expect("failed to bind server address");

    println!("listening on http://{addr}");
    axum::serve(listener, app)
        .await
        .expect("server exited unexpectedly");
}

async fn root() -> &'static str {
    "Rust Axum + PostgreSQL"
}

async fn health() -> Json<Value> {
    Json(json!({ "status": "ok" }))
}
