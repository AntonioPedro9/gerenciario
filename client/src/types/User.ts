export type ICreateUserRequest = {
  name: string;
  email: string;
  password: string;
};

export type ILoginUserRequest = {
  email: string;
  password: string;
};
