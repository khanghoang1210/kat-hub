import React, { useEffect } from "react";
import SignIn from "./SignIn";
import SignUp from "./SignUp";
import { useRecoilValue, useSetRecoilState } from "recoil";
import { authModalState } from "@/atoms/authModalAtom";

type AuthModalProps = {};

const AuthModal: React.FC<AuthModalProps> = () => {
  const authModal = useRecoilValue(authModalState)

  return (
    <>
      <div className="flex items-center justify-center pt-5 pb-20">
        <div className="rounded-lg shadow bg-white-900 flex items-center justify-center content-center w-[384px]">
        {authModal.type ==='login' ? <SignIn/> :authModal.type === 'register'? <SignUp/> : null}
        </div>
      </div>
    </>
  );
};
export default AuthModal;
