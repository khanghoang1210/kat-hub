import { authModalState } from '@/atoms/authModalAtom';
import React, { useEffect } from 'react';
import { useRecoilState, useRecoilValue } from 'recoil';
import Image from 'next/image';
import AuthModal from '@/components/auth/AuthModal';
import { useRouter } from 'next/router';

type ResetPassProps = {
    
};

const ResetPass:React.FC<ResetPassProps> = () => {
    const [authModal, setAuthModalState] = useRecoilState(authModalState);
    const router = useRouter();
    useEffect(() => {
        if (router.pathname === '/recover-password') {
            setAuthModalState((prev) => ({
                ...prev,
                type: 'resetPassword'
            }));
        }
    }, [router.pathname, setAuthModalState]);
    
   
      return (
        <div className=" flex flex-col items-center justify-center">
          <div className="flex justify-center items-center mt-2">
            <Image src="/Logomark.svg" alt="Logo" height={40} width={40} />
            <h1 className="text-lg font-bold ml-2">Kathub</h1>
          </div>
          { <AuthModal/>} 
        </div>
      );
}
export default ResetPass;