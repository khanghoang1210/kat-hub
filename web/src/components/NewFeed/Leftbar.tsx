import Image from "next/image";
import React from "react";

type LeftbarProps = {};

const Leftbar: React.FC<LeftbarProps> = () => {
  return (
    <div className="absolute rounded-lg left-20 top-20 w-64 h-auto p-0 mt-2 bg-white-900 shadow text-black-600">
      <div className="relative  w-full ">
        <img
          src="https://via.placeholder.com/256x100" // Replace with the actual cover image source
          alt="Cover"
          className="w-screen h-fit object-cover rounded-t-lg"
        />
        <div className="absolute top-16 left-4 ">
        <Image
          src={"/assets/icons/user.svg"}
          alt="avt"
          width={60}
          height={60}
          className=" rounded-full border-2 border-black-300"
        />
        </div>
        <div className="mt-10 ml-4">
          <h2 className="font-bold text-lg">Robert Fox</h2>
          <p className="text-gray-500">Software Engineer</p>
        </div>
      </div>
      <div className="mt-8 space-y-4 m-3">
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
