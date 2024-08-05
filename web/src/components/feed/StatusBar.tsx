import Image from "next/image";
import React from "react";

type StatusBarProps = {
  avatarUrl: string;
};

const StatusBar: React.FC<StatusBarProps> = ({avatarUrl}) => {
  return (
    <div className=" left-[365px] w-5/12 h-28  rounded-lg absolute bg-white-900 p-4 mt-4 shadow text-black-600">
      <div className="flex items-center">
        <Image
          src={avatarUrl ? avatarUrl : "/assets/icons/user.svg" }
          alt="avt"
          width={40}
          height={40}
          className=" rounded-full border-1 border-black-300"
        />
        <input
          className="outline-none absolute border-b-2 border-white-400 top-3 right-8 w-4/5 h-12 text-black-300"
          type="text"
          placeholder="What's on your mind?"
        />
      

      <button className="rounded-2xl absolute right-8 bg-blue-900 w-16 h-8 mt-24 text-white-900">
        Post
      </button>
      </div>
      <button className="absolute flex left-16 top-20 text-sm w-32 font-medium text-black-800">
        <Image
          src={"/assets/icons/Media.svg"}
          alt="media"
          width={25}
          height={20}
        />
        <p className="absolute right-6 top-0">Add Media</p>
      </button>
    </div>
  );
};
export default StatusBar;
