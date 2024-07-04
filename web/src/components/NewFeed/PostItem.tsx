import Image from "next/image";
import React from "react";

type PostItemProps = {
  username: string;
  userTitle: string;
  postTime: string;
  postContent: string;
  userImage: string;
};

const PostItem: React.FC<PostItemProps> = ({
  username,
  userTitle,
  postTime,
  postContent,
  userImage,
}) => {
  return (
    <div className="p-3 bg-white-900 rounded-lg shadow w-5/12 mx-80 top-32 relative mt-4">
      <div className="flex items-center mb-2 h-11">
        <Image
          className="w-10 h-10 rounded-full border-2 border-black-300"
          src={userImage}
          alt={username}
          width={20}
          height={20}
        
        />
        <div className="ml-3">
          <h2 className="text-sm font-bold">{username}</h2>
          <p className="text-xs text-gray-600">{userTitle}</p>
        </div>
        <span className="ml-auto text-xs text-gray-500">{postTime}</span>
      </div>
      <div className="relative flex py-1 items-center w-full">
        <div className="flex-grow border-t border-white-200"></div>
      </div>

      <p className="mb-4 text-black-800 font-normal">{postContent}</p>
      <div className="flex items-center space-x-4">
        <button className="flex items-center text-gray-500 hover:text-gray-700 text-black-500 font-medium">
         <Image
         src={"/assets/icons/Comment.svg"}
         alt="comment"
         width={20}
         height={20}
         />
         
          <p className="ml-2 -mt-1">Comment</p>
        </button>
        <button className="flex hover:text-gray-700 -mt-1 pl-96">
         <Image src={'/assets/icons/like.svg'} width={20} height={20} alt="like"/>
       
        </button>
      </div>
    </div>
  );
};

export default PostItem;
