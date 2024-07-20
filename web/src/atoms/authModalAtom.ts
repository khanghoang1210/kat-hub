import { atom } from "recoil";

type AuthModalState ={
    isOpen: boolean;
    type: 'login'|'register'|'forgotPassword'|'resetPassword'
}

const initialAuthModalState:AuthModalState = {
    isOpen: true,
    type: 'login'
};

export const authModalState = atom<AuthModalState>({
    key: "authModalState",
    default: initialAuthModalState,
   
});