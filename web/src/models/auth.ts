
export type LoginReq = {
    userName: string;
    password: string;
}

export type SignUpReq = {
    userName: string;
    email: string;
    fullName: string;
    password: string;
    confirmPassword: string;
}