import * as yup from "yup";

const clientSchema = yup.object().shape({
  cpf: yup
    .string()
    .matches(/^\d{11}$/, "CPF deve ter 11 dígitos")
    .required("Preencha todos os campos obrigatórios"),

  name: yup.string().required("Preencha todos os campos obrigatórios"),

  email: yup.string().email("Email inválido"),

  phone: yup
    .string()
    .matches(/^\d{10,11}$/, "Telefone deve ter 10 ou 11 dígitos")
    .required("Preencha todos os campos obrigatórios"),
});

const serviceSchema = yup.object().shape({
  name: yup.string().required("Preencha todos os campos obrigatórios"),

  description: yup.string(),

  duration: yup.number().integer().min(0, "Duração deve ser um número inteiro maior ou igual a 0"),

  price: yup.number().min(0, "O preço deve ser maior ou igual a zero"),
});

export { clientSchema, serviceSchema };
