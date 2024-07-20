import Image from "next/image";
import React from "react";

type AddFriendProps = {
  userName: string;
  avatar: string;
  userTitle: string;
};

const AddFriend: React.FC<AddFriendProps> = (friend: AddFriendProps) => {
  return (
    <div className="inline-flex p-1 mt-2 w-full h-14">
      <Image
        className="w-10 h-10 rounded-full border-2 border-black-300"
        src={friend.avatar}
        alt={friend.userName}
        width={20}
        height={20}
      />
      <div className="ml-3">
        <h2 className="text-sm font-medium text-black-900">
          {friend.userName}
        </h2>
        <p className="text-xs text-black-800">{friend.userTitle}</p>
      </div>
      <button className="w-8 h-8 pl-[210px] cursor-pointer absolute">
        <svg
          width="32"
          height="32"
          viewBox="0 0 32 32"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <rect width="32" height="32" rx="6" fill="#F1F4F9" />
          <path
            d="M16.4167 11V21"
            stroke="#27364B"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
          <path
            d="M11.4167 16H21.4167"
            stroke="#27364B"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
      </button>
    </div>
  );
};
export default AddFriend;
