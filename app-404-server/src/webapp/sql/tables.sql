
-- create database kone_guest;
create table `guest` (
	`guest_id` int NOT NULL AUTO_INCREMENT,
	`name` varchar(64) NOT NULL,
	`create_time` datetime,
	`visit_date` datetime,
	`face_id` varchar(64),
	`visit_building` varchar(128),
	`visit_floor` varchar(64),
	`visit_room` varchar(64),
	PRIMARY KEY (`guest_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table `visiting_guest` (
	`guest_id` int NOT NULL,
	`type` int NOT NULL DEFAULT 0,
	`is_visiting` int NOT NULL DEFAULT 0,
	`last_visit_time` datetime,
	`position_info` text, 
	PRIMARY KEY (`guest_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table `staff` (
	`user_id` int NOT NULL,
	`type` varchar(64),
	`name` int NOT NULL DEFAULT 0,
	`face_id` varchar(64),
	PRIMARY KEY (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

