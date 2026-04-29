use axum::{routing::get, Json, Router};
use db::db_health;
use serde_json::{json, Value};
use std::{env, net::SocketAddr};

mod db;

#[tokio::main]
async fn main() {
    let database_url = env::var("DATABASE_URL")
        .unwrap_or_else(|_| "postgres://postgres:postgres@localhost:5432/app".to_string());

    let app = Router::new()
        .route("/", get(root))
        .route("/health", get(health))
        .route("/db-health", get(move || db_health(database_url.clone())));

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
