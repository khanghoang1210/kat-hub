
import { authModalState } from '@/atoms/authModalAtom';
import AuthModal from '@/components/auth/AuthModal';
import Image from 'next/image';
import React from 'react';
import { useRecoilValue } from 'recoil';

type  AuthPageProps = {
    
};

const  AuthPage:React.FC< AuthPageProps> = () => {
    
  const authModal = useRecoilValue(authModalState)
    return (
        <div className=" flex flex-col items-center justify-center">
          <div className="flex justify-center items-center mt-2">
            <Image src="/Logomark.svg" alt="Logo" height={40} width={40} />
            <h1 className="text-lg font-bold ml-2">Kathub</h1>
          </div>
          {authModal.isOpen && <AuthModal/>} 
        </div>
      );
}
export default AuthPage;