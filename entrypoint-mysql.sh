echo "MySQL Host: $MYSQL_HOST"
echo "MySQL Port: $MYSQL_PORT"
echo "MySQL Database: $MYSQL_DATABASE"

#until mysqladmin ping -h"$MYSQL_HOST" -P"$MYSQL_PORT" -u"root" -p"$MYSQL_ROOT_PASSWORD" --silent; do
#    sleep 1
#done

mysql -h"$MYSQL_HOST" -P"$MYSQL_PORT" -u"root" -p"$MYSQL_ROOT_PASSWORD" "$MYSQL_DATABASE" < /app/create-tables.sql