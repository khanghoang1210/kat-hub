import Image from "next/image";
import React from "react";

type StatusBarProps = {};

const StatusBar: React.FC<StatusBarProps> = () => {
  return (
    <div className="ms-80 w-5/12 h-36 rounded-lg absolute bg-white-900 h-36 p-4 mt-1 ml-8 shadow-md text-black-600">
      <div className="h-2/3 p-0">
        <Image
          src={"/assets/icons/user.svg"}
          alt="avt"
          width={40}
          height={40}
          className=" rounded-full mt-4 border-2 border-black-300"
        />
        <input
          className="outline-none absolute border-b-2 border-white-400 sm:text-sm top-8 right-8 w-4/5 h-12 text-black-400"
          type="text"
          placeholder="Bạn đang nghĩ gì?"
        />
        <button className="rounded-2xl absolute right-8 bg-blue-900 w-16 h-8 mt-4 text-white-900">
          Post
        </button>
        <div className="mt-6 ml-2">
        <button className="absolute left-16 text-sm w-32 font-medium text-black-800">
        <Image
            src={"/assets/icons/Media.svg"}
            alt="media"
            width={25}
            height={20}
        
          />
          <p className="absolute right-6 top-0">
          Add Media
            </p> 
          </button>
          
            
        </div>
      </div>
    </div>
  );
};
export default StatusBar;
