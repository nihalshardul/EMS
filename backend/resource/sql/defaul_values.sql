USE EMS;
/* user_info values*/

insert into user_info (id,first_name,last_name,email,password,is_user_root) values(1,'ACL','SYSTEM','acl.system@acl.com','password',1);
insert into user_info (id,first_name,last_name,email,password,is_user_root) values(420,'Tushar','Pathare','tushar.pathare@gmail.com','Tushar@123',1);
/* groups*/
insert into groups(id,group_name) values(1,'root');

/*permission table */
insert into permission(id,permission_info) values(100,"read");
insert into permission(id,permission_info) values(010,"write");
insert into permission(id,permission_info) values(001,"execute");
insert into permission(id,permission_info) values(110,"read write");
insert into permission(id,permission_info) values(101,"read excute");
insert into permission(id,permission_info) values(111,"read write excute");



/* resources types*/
insert into resources_type values(1,"File");
insert into resources_type values(2,"Folder");

/* resources*/
insert into resources(id,resource_type_id,resource_name,resource_path) values(1,2,'ACL_STORAGE','/home/Nihal/ACL');

insert into user_group (id,user_id) values(1,1);

insert into user_permission (resource_id,id,permission_id) values(1,1,101);

insert into group_permission (resource_id,id,permission_id) values(1,1,100);

