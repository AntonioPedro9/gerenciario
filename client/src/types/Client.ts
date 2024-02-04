export type IClient = {
  id: number;
  cpf: string;
  name: string;
  email: string;
  phone: string;
  userID: string;
};

export type ICreateClientRequest = {
  cpf: string;
  name: string;
  email?: string;
  phone: string;
  userID: string;
};

export type IUpdateClientRequest = {
  id: number;
  cpf?: string | null;
  name?: string | null;
  email?: string | null;
  phone?: string | null;
  userID: string;
};
