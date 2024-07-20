import { authModalState } from '@/atoms/authModalAtom';
import React from 'react';
import { useSetRecoilState } from 'recoil';

type SignInProps = {
    
};

const SignIn:React.FC<SignInProps> = () => {
    const setAuthModalState = useSetRecoilState(authModalState);
    const handleClick = (type:"login"|"register"|"forgotPassword") =>{
        setAuthModalState((prev) =>({...prev, type}));
    }
    
    return <form className=' space-y-6 px-7 pt-10 pb-4  w-[384px] h-[350px]'>

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
    <div>
       
        <input
     
            type='password'
            name='password'
            id='password'
            className='
    border-2 border-white-400 outline-none sm:text-sm rounded-lg block w-full p-2.5
    bg-gray-600 text-white
'
            placeholder='Password'
        />
    </div>

    <button
        type='submit'
        className='w-full text-white focus:ring-blue-300 font-medium rounded-lg
        text-sm text-white-900 px-5 py-2.5 text-center bg-black-900 hover:bg-black-800
    '
    >
        Login
    </button>
    <button className='flex w-full justify-end'>
        <a href='#' className='text-sm block text-black-500 hover:underline w-full text-right'>
            Forgot Password?
        </a>
    </button>
    <div className='text-sm text-black-600'>
        Don&apos;t have an account? &nbsp;{" "}
        <a href='#' className='text-black-800 hover:underline font-medium' onClick={()=>handleClick("register")}>
           Sign Up
        </a>
    </div>
</form>
}
export default SignIn;