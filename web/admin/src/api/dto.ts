/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface DtoCreateCategoryRequest {
  color: string;
  name: string;
  parentId?: string;
}

export interface DtoCreateCategoryResponse {
  color?: string;
  createdAt?: string;
  id?: string;
  name?: string;
  parentId?: string;
  updatedAt?: string;
}

export interface DtoCreateLabelRequest {
  color: string;
  name: string;
}

export interface DtoCreateLabelResponse {
  color?: string;
  createdAt?: string;
  id?: string;
  name?: string;
  updatedAt?: string;
}

export interface DtoCreateUserRequest {
  code: string;
  email: string;
  nickname?: string;
  password: string;
  username: string;
}

export interface DtoCreateUserResponse {
  createdAt?: string;
  email?: string;
  id?: string;
  nickname?: string;
  updatedAt?: string;
  username?: string;
}

export interface DtoListResponseDtoCreateCategoryResponse {
  count?: number;
  list?: DtoCreateCategoryResponse[];
}

export interface DtoListResponseDtoCreateLabelResponse {
  count?: number;
  list?: DtoCreateLabelResponse[];
}

export interface DtoListResponseDtoCreateUserResponse {
  count?: number;
  list?: DtoCreateUserResponse[];
}

export interface DtoLoginRequest {
  password: string;
  username: string;
}

export interface DtoModifyCategoryRequest {
  color?: string;
  id: string;
  name?: string;
  parentId?: string;
}

export interface DtoModifyLabelRequest {
  color?: string;
  id: string;
  name?: string;
}

export interface DtoModifyUserRequest {
  avatar?: string;
  id: string;
  nickname?: string;
  password?: string;
}

export interface DtoSuccessEmptyResponse {
  /** @example 0 */
  code?: number;
  /** @example "" */
  data?: string;
  /** @example "success" */
  msg?: string;
}

export interface ResponseBodyDtoCreateCategoryResponse {
  code?: number;
  data?: DtoCreateCategoryResponse;
  msg?: string;
}

export interface ResponseBodyDtoCreateLabelResponse {
  code?: number;
  data?: DtoCreateLabelResponse;
  msg?: string;
}

export interface ResponseBodyDtoCreateUserResponse {
  code?: number;
  data?: DtoCreateUserResponse;
  msg?: string;
}

export interface ResponseBodyDtoListResponseDtoCreateCategoryResponse {
  code?: number;
  data?: DtoListResponseDtoCreateCategoryResponse;
  msg?: string;
}

export interface ResponseBodyDtoListResponseDtoCreateLabelResponse {
  code?: number;
  data?: DtoListResponseDtoCreateLabelResponse;
  msg?: string;
}

export interface ResponseBodyDtoListResponseDtoCreateUserResponse {
  code?: number;
  data?: DtoListResponseDtoCreateUserResponse;
  msg?: string;
}

export interface ResponseBodyString {
  code?: number;
  data?: string;
  msg?: string;
}
