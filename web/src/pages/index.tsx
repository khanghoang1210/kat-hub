import Image from "next/image";
import { Inter } from "next/font/google";
import Topbar from "@/components/Topbar/Topbar";
import Leftbar from "@/components/Leftbar/Leftbar";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  return (
    <div>
      <Topbar/>
      <Leftbar/>
    </div>
  );
}
