mod todo;


pub fn add(left: u64, right: u64) -> u64 {
    left + right
}

#[cfg(test)]
mod tests {
    use super::*;
    // use sea_orm::{entity::*, query::*, sea_query::*, DbConn, DbErr, SeaOrm};
    use tokio;

    // async fn setup_db() -> DbConn {
    //     let database_url = "mysql://root:password@localhost:3306/todolist";
    //     let db = DbConn::connect(database_url).await.unwrap();
    //     db
    // }



    #[tokio::test]
    async fn it_works() {
        let result = add(2, 2);
        assert_eq!(result, 4);
    }
}
