import Image from "next/image";
import { Inter } from "next/font/google";
import Topbar from "@/components/Topbar/Topbar";
import Leftbar from "@/components/Leftbar/Leftbar";
import StatusBar from "@/components/NewFeed/StatusBar";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  return (
    <div>
      <Topbar/>
      <StatusBar/>
      <Leftbar/>
    </div>
  );
}
