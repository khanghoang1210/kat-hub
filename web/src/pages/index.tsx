import Image from "next/image";
import { Inter } from "next/font/google";
import Topbar from "@/components/Topbar/Topbar";
import Leftbar from "@/components/Leftbar/Leftbar";
import StatusBar from "@/components/NewFeed/StatusBar";
import PostItem from "@/components/NewFeed/PostItem";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  return (
    <>
      <div>
        <Topbar />
        <div className="flex">
          <Leftbar />
          <div className="flex-grow">
            <StatusBar />
            <div className="">
              <PostItem
                username="Bessie Cooper"
                userTitle="Digital Marketer"
                postTime="7 hours ago"
                postContent="In today's fast-paced, digitally driven world, digital marketing is not just a strategy. it's a necessity for businesses of all sizes. 🚀"
                userImage="/assets/icons/user.svg"
              />
               <PostItem
                username="Bessie Cooper"
                userTitle="Digital Marketer"
                postTime="7 hours ago"
                postContent="In today's fast-paced, digitally driven world, digital marketing is not just a strategy. it's a necessity for businesses of all sizes. 🚀"
                userImage="/assets/icons/user.svg"
              />
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
