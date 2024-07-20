import { authModalState } from "@/atoms/authModalAtom";
import React from "react";
import { useSetRecoilState } from "recoil";

type SignUpProps = {};

const SignUp: React.FC<SignUpProps> = () => {

    const setAuthModalState = useSetRecoilState(authModalState);
    const handleClick = (type:"login"|"register"|"forgotPassword") =>{
        setAuthModalState((prev) =>({...prev, type}));
    }

  return (
    <form className=" space-y-6 px-7 pt-7 pb-4  w-[384px] h-[480px]">
      <div>
        <input
          type="userName"
          name="userName"
          id="userName"
          className="
    border-2 outline-none sm:text-sm rounded-lg block w-full p-2.5
    bg-gray-600 border-white-400 placeholder-gray-400 text-white
"
          placeholder="User Name"
        />
      </div>
      <div>
        <input
          type="email"
          name="email"
          id="email"
          className="
    border-2 outline-none sm:text-sm rounded-lg block w-full p-2.5
    bg-gray-600 border-white-400 placeholder-gray-400 text-white
"
          placeholder="Email"
        />
      </div>
      <div>
        <input
          type="fullName"
          name="fullName"
          id="fullName"
          className="
    border-2 outline-none sm:text-sm rounded-lg block w-full p-2.5
    bg-gray-600 border-white-400 placeholder-gray-400 text-white
"
          placeholder="Full Name"
        />
      </div>
      <div>
        <input
          type="password"
          name="password"
          id="password"
          className="
    border-2 border-white-400 outline-none sm:text-sm rounded-lg block w-full p-2.5
    bg-gray-600 text-white
"
          placeholder="Password"
        />
      </div>
      <div>
        <input
          type="confirmPassword"
          name="confirmPassword"
          id="confirmPassword"
          className="
    border-2 border-white-400 outline-none sm:text-sm rounded-lg block w-full p-2.5
    bg-gray-600 text-white
"
          placeholder="Confirm Password"
        />
      </div>

      <button
        type="submit"
        className="w-full text-white focus:ring-blue-300 font-medium rounded-lg
        text-sm text-white-900 px-5 py-2.5  text-center bg-black-900 hover:bg-black-800
    "
      >
        Continue
      </button>
      <div className="text-sm text-black-600 justify-center flex relative -mt-1">
        Have an account? &nbsp;{" "}
        <a href="#" className="text-black-800 hover:underline font-medium"  onClick={()=>handleClick("login")}>
          Log In
        </a>
      </div>
    </form>
  );
};
export default SignUp;
