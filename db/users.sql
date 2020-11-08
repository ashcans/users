-- ----------------------------
-- Sequence structure for role_role_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."role_role_id_seq";
CREATE SEQUENCE "public"."role_role_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_role_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."user_role_id_seq";
CREATE SEQUENCE "public"."user_role_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_token_user_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."user_token_user_id_seq";
CREATE SEQUENCE "public"."user_token_user_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_user_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."user_user_id_seq";
CREATE SEQUENCE "public"."user_user_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for menu_menu_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."menu_menu_id_seq";
CREATE SEQUENCE "public"."menu_menu_id_seq"
    INCREMENT 1
    MINVALUE  1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;

-- ----------------------------
-- Sequence structure for role_menu_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."role_menu_id_seq";
CREATE SEQUENCE "public"."role_menu_id_seq"
    INCREMENT 1
    MINVALUE  1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS "public"."role";
CREATE TABLE "public"."role" (
  "role_id" int8 NOT NULL DEFAULT nextval('role_role_id_seq'::regclass),
  "role_name" varchar(100) COLLATE "pg_catalog"."default",
  "remark" varchar(100) COLLATE "pg_catalog"."default",
  "create_user_id" int8,
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6),
  CONSTRAINT role_pkey PRIMARY KEY (role_id)
)
;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "public"."user";
CREATE TABLE "public"."user" (
  "user_id" int8 NOT NULL DEFAULT nextval('user_user_id_seq'::regclass),
  "username" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(100) COLLATE "pg_catalog"."default",
  "salt" varchar(20) COLLATE "pg_catalog"."default",
  "nick_name" varchar(100) COLLATE "pg_catalog"."default",
  "email" varchar(100) COLLATE "pg_catalog"."default",
  "mobile" varchar(100) COLLATE "pg_catalog"."default",
  "status" int4,
  "create_user_id" int8,
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6),
  CONSTRAINT user_pkey PRIMARY KEY (user_id)
)
;

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_role";
CREATE TABLE "public"."user_role" (
  "id" int8 NOT NULL DEFAULT nextval('user_role_id_seq'::regclass),
  "user_id" int8,
  "role_id" int8,
  CONSTRAINT user_role_pkey PRIMARY KEY (user_id)
)
;

-- ----------------------------
-- Table structure for user_token
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_token";
CREATE TABLE "public"."user_token" (
  "user_id" int8 NOT NULL DEFAULT nextval('user_token_user_id_seq'::regclass),
  "token" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "expire_time" timestamp(6),
  "update_time" timestamp(6),
  CONSTRAINT user_token_pkey PRIMARY KEY (user_id),
  CONSTRAINT user_token_token_key UNIQUE (token)
)
;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS "public"."menu";
CREATE TABLE public.menu
(
    "menu_id" bigint NOT NULL DEFAULT nextval('menu_menu_id_seq'::regclass),
    "parent_id" bigint,
    "name" character varying(50) COLLATE pg_catalog."default",
    "url" character varying(200) COLLATE pg_catalog."default",
    "perms" character varying(500) COLLATE pg_catalog."default",
    "type" integer,
    "icon" character varying(50) COLLATE pg_catalog."default",
    "order_num" integer,
    CONSTRAINT menu_pkey PRIMARY KEY (menu_id)
)
;

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
CREATE TABLE public.role_menu
(
    "id" bigint NOT NULL DEFAULT nextval('role_menu_id_seq'::regclass),
    "role_id" bigint,
    "menu_id" bigint,
    CONSTRAINT role_menu_pkey PRIMARY KEY (id)
)
;