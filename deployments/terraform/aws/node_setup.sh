#!/bin/bash -ex

cd "$(dirname "$0")"

yum update -y
yum -y install httpd httpd-tools

service httpd start
chkconfig httpd on
usermod -a -G apache ec2-user
chown -R ec2-user:apache /var/www
chmod 2775 /var/www

find /var/www -type d -exec chmod 2775 {} \;
find /var/www -type f -exec chmod 0664 {} \;

cd /var/www/html

echo '<html><body>hello</body></html>' > index.html
