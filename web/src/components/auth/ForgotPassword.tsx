import { authModalState } from '@/atoms/authModalAtom';
import React from 'react';
import { useSetRecoilState } from 'recoil';

type ForgotPasswordProps = {
    
};

const ForgotPassword:React.FC<ForgotPasswordProps> = () => {
    const setAuthModalState = useSetRecoilState(authModalState);
    const handleClick = (type:"login"|"register"|"forgotPassword") =>{
        setAuthModalState((prev) =>({...prev, type}));
    }
    
    return <form className=' space-y-6 px-7 pt-10 pb-4  w-[384px] h-[350px]'>
        <p className='justify-center flex text-black-900 font-medium text-base'>Forgot Password</p>
        <p className='justify-center items-center text-center flex text-black-500 font-normal text-base'>Enter your email to reset your password and access your account.</p>

    <div>
        <input
            type='userName'
            name='userName'
            id='userName'
            className='
    border-2 outline-none sm:text-sm rounded-lg block w-full p-2.5
    bg-gray-600 border-white-400 placeholder-gray-400 text-white
'
            placeholder='Email'
        />
    </div>
  
    <button
        type='submit'
        className='w-full text-white focus:ring-blue-300 font-medium rounded-lg
        text-sm text-white-900 px-5 py-2.5 text-center bg-black-900 hover:bg-black-800
    '
    >
        Send Reset Link
    </button>
   
</form>
}
export default ForgotPassword;