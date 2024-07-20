
import AuthModal from '@/components/auth/AuthModal';
import Image from 'next/image';
import React from 'react';

type  AuthPageProps = {
    
};

const  AuthPage:React.FC< AuthPageProps> = () => {
    
    return (
        <div className="h-screen flex flex-col items-center justify-center">
          <div className="flex items-center mt-12">
            <Image src="/Logomark.svg" alt="Logo" height={40} width={40} />
            <h1 className="text-lg font-bold ml-2">Kathub</h1>
          </div>
          <AuthModal />
        </div>
      );
}
export default AuthPage;