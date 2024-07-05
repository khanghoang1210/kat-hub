import Image from "next/image";
import { Inter } from "next/font/google";
import Topbar from "@/components/Topbar/Topbar";
import Leftbar from "@/components/NewFeed/Leftbar";
import StatusBar from "@/components/NewFeed/StatusBar";
import PostItem from "@/components/NewFeed/PostItem";
import FriendItem from "@/components/NewFeed/FriendItem";

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
            <FriendItem/>
            <div className="">
              <PostItem
                username="Bessie Cooper"
                userTitle="Digital Marketer"
                postTime="7 hours ago"
                postContent="In today's as sjhd jashd cija aksdj casi caksji laksdj caslkdji alksdj cajsjhdl kajsd ca asd j jas jhds i al as d ask kasd klkj l kj iij nm mnsd jd la p qe cnas asdj sad kh kj kj kj kj kj kj kj kj k jkj kj kj k jk j kj kj k kj jj jj jj jj jj jj jj jj jj jj j jj j j j j jj j jj jj j j jj jj jj jj jj jj j jj j jj j j j j j jj fast-paced, digitally driven world, digital marketing is not just a strategy. it's a necessity for businesses of all sizes. 🚀"
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
