create table posts(
id integer primary key,
title text ,
content text,
user_id integer references users(id),
tags text[]

)