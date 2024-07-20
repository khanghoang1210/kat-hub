import { authModalState } from '@/atoms/authModalAtom';
import React from 'react';
import { useSetRecoilState } from 'recoil';

type ResetPasswordProps = {
    
};

const ResetPassword:React.FC<ResetPasswordProps> = () => {
    const setAuthModalState = useSetRecoilState(authModalState);
    const handleClick = (type:"login"|"register"|"forgotPassword") =>{
        setAuthModalState((prev) =>({...prev, type}));
    }
    
    return <form className=' space-y-6 px-7 pt-10 pb-4  w-[384px] h-[350px]'>
 <p className='justify-center flex text-black-900 font-medium text-base pb-10'>Reset Password</p>
    <div>
        <input
            type='password'
            name='password'
            id='password'
            className='
    border-2 outline-none sm:text-sm rounded-lg block w-full p-2.5
    bg-gray-600 border-white-400 placeholder-gray-400 text-white
'
            placeholder='New Password'
        />
    </div>
    <div>
       
        <input
     
            type='confirmPassword'
            name='confirmPassword'
            id='confirmPassword'
            className='
    border-2 border-white-400 outline-none sm:text-sm rounded-lg block w-full p-2.5
    bg-gray-600 text-white
'
            placeholder='Confirm Password'
        />
    </div>

    <button
        type='submit'
        className='w-full text-white focus:ring-blue-300 font-medium rounded-lg
        text-sm text-white-900 px-5 py-2.5 text-center bg-black-900 hover:bg-black-800
    '
    >
        Reset
    </button>
  
</form>
}
export default ResetPassword;