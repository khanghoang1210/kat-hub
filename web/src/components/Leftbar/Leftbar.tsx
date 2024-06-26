import Image from "next/image";
import React from "react";

type LeftbarProps = {};

const Leftbar: React.FC<LeftbarProps> = () => {
  return (
    <div className="rounded-lg relative bg-white-900 w-64 h-auto p-4 mt-1 ml-8 shadow-md text-black-600">
      <div className="relative  w-full">
        <img
          src="https://via.placeholder.com/256x100" // Replace with the actual cover image source
          alt="Cover"
          className="w- h-24 object-cover"
        />
        <div className="absolute top-16 left-4">
          <img
            src="https://via.placeholder.com/50" // Replace with the actual avatar image source
            alt="User Avatar"
            className="w-16 h-16 rounded-full border-4 border-white"
          />
        </div>
        <div className="mt-8 ml-4">
          <h2 className="font-bold text-lg">Robert Fox</h2>
          <p className="text-gray-500">Software Engineer</p>
        </div>
      </div>
      <div className="mt-8 space-y-4">
        <button className="flex items-center space-x-2 p-1.5 w-full text-left hover:bg-gray-100 rounded">
          <Image
            src={"/assets/icons/Home.svg"}
            width={20}
            height={20}
            alt="home"
          />
          <span>Home</span>
        </button>
        <button className="flex items-center space-x-2 p-1.5 w-full text-left hover:bg-gray-100 rounded">
          <Image
            src={"/assets/icons/user.svg"}
            width={20}
            height={20}
            alt="profile"
          />
          <span>Profile</span>
        </button>
        <button className="flex items-center space-x-2 p-1.5 w-full text-left hover:bg-gray-100 rounded">
          <Image
            src={"/assets/icons/Send.svg"}
            width={20}
            height={20}
            alt="send"
          />
          <span>Messages</span>
        </button>
        <button className="flex items-center space-x-2 p-1.5 w-full text-left hover:bg-gray-100 rounded">
          <Image
            src={"/assets/icons/Notification.svg"}
            width={20}
            height={20}
            alt="Noti"
          />
          <span>Notifications</span>
        </button>
      </div>
    </div>
  );
};

export default Leftbar;
