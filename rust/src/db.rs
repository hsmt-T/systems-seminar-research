use axum::{extract::State, Json};
use serde_json::{json, Value};
use sqlx::{postgres::PgPoolOptions, PgPool};

use crate::AppState;

pub type DbPool = PgPool;

pub async fn create_pool(database_url: &str) -> Result<DbPool, sqlx::Error> {
    PgPoolOptions::new()
        .max_connections(5)
        .connect(database_url)
        .await
}

pub async fn db_health(State(state): State<AppState>) -> Json<Value> {
    match sqlx::query_scalar::<_, i32>("select 1")
        .fetch_one(&state.db_pool)
        .await
    {
        Ok(value) => Json(json!({ "status": "ok", "select": value })),
        Err(error) => Json(json!({ "status": "error", "message": error.to_string() })),
    }
}