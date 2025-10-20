create database "GolangFirebase";
create schema golangfirebase;
alter database "GolangFirebase" set search_path to golangfirebase;
create extension if not exists pgcrypto;