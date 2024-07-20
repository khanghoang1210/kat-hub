import React from 'react';
import AddFriend from './AddFriend';

type FriendItemProps = {
    
};

const FriendItem:React.FC<FriendItemProps> = () => {
    
    return <div className='float-end mr-16 w-[256px] h-[400px] mt-4 rounded-lg bg-white-900 '>
        <p className='pl-8 pt-2 text-base font-medium text-black-900'>Suggested Friends</p>
        <div className="relative flex py-4 items-center w-full">
        <div className="flex-grow border-t border-white-200"></div>
        
      </div>
      <AddFriend userName='Olivia Anderson' avatar='/assets/icons/user.svg' userTitle='Financial Analyst'/>
      <AddFriend userName='Thomas Baker' avatar='/assets/icons/user.svg' userTitle='Project Manager'/>
      <AddFriend userName='Lily Lee' avatar='/assets/icons/user.svg' userTitle='Graphic Designer'/>
      <AddFriend userName='Lily Lee' avatar='/assets/icons/user.svg' userTitle='Graphic Designer'/>
    </div>
}
export default FriendItem;