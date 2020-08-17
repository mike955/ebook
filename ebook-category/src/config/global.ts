module.exports = {
  mysqlHost: process.env.MYSQL_HOST || "192.168.56.101",
  mysqlPort: process.env.MYSQL_PORT || 3306,
  mysqlUser: process.env.MYSQL_USER || 'root',
  mysqlPassword: process.env.MYSQL_PASSWORD || '@Cai3564423',
  mysqlDatabase: process.env.MYSQL_DATABASE || 'ebook',
  mysqlLogging: process.env.MYSQL_HOST || true,
}