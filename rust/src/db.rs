use axum::Json;
use serde_json::{json, Value};
use tokio_postgres::NoTls;

pub async fn db_health(database_url: String) -> Json<Value> {
    match tokio_postgres::connect(&database_url, NoTls).await {
        Ok((client, connection)) => {
            tokio::spawn(async move {
                if let Err(error) = connection.await {
                    eprintln!("postgres connection error: {error}");
                }
            });

            match client.query_one("select 1", &[]).await {
                Ok(row) => {
                    let value: i32 = row.get(0);
                    Json(json!({ "status": "ok", "select": value }))
                }
                Err(error) => Json(json!({ "status": "error", "message": error.to_string() })),
            }
        }
        Err(error) => Json(json!({ "status": "error", "message": error.to_string() })),
    }
}
