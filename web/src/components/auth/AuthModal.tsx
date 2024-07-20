import React from "react";
import SignUp from "./SignIn";

type AuthModalProps = {};

const AuthModal: React.FC<AuthModalProps> = () => {
  return (
    <>
      <div className="flex items-center justify-center pt-10 pb-20">
        <div className="rounded-md shadow bg-white-900 flex items-center justify-center content-center w-[384px]">
          <SignUp />
        </div>
      </div>
    </>
  );
};
export default AuthModal;
