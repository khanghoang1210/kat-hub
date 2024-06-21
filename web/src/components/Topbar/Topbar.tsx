import React, { useRef } from "react";

type TopbarProps = {};

const Topbar: React.FC<TopbarProps> = () => {
  const svgRef = useRef<SVGSVGElement>(null);
  const inputRef = useRef<HTMLInputElement>(null);

  return (
    <nav className="relative flex h-[50px] w-full shrink-0 items-center px-5 bg-dark-layer-1 text-dark-gray-7">
      <div className="relative w-1/2 mx-auto mt-10">
        <svg
          ref={svgRef}
          className="absolute top-1/2 left-4 transform -translate-y-1/2 pointer-events-none transition-opacity duration-0 ease-in-out"
          width="17"
          height="17"
          viewBox="0 0 20 20"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M9.16667 15.8333C12.8486 15.8333 15.8333 12.8486 15.8333 9.16667C15.8333 5.48477 12.8486 2.5 9.16667 2.5C5.48477 2.5 2.5 5.48477 2.5 9.16667C2.5 12.8486 5.48477 15.8333 9.16667 15.8333Z"
            stroke="#5D6778"
            strokeWidth="1.75"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
          <path
            d="M17.5 17.5L13.875 13.875"
            stroke="#5D6778"
            strokeWidth="1.75"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
        </svg>

        <input
          ref={inputRef}
          className="border-2 border-blue-100 sm:text-sm rounded-md w-full h-9 text-black pl-10 pr-4 transition-all duration-300 ease-in-out focus:border-blue-300 outline-none hover:border-blue-300 "
          type="text"
          placeholder="Search"
          onFocus={() => {
            if (svgRef.current) {
              svgRef.current.style.opacity = '0';
            }
            if (inputRef.current) {
              inputRef.current.style.paddingLeft = '0.75rem';
            }
          }}
          onBlur={() => {
            if (svgRef.current) {
              svgRef.current.style.opacity = '1';
            }
            if (inputRef.current) {
              inputRef.current.style.paddingLeft = '2.5rem';
            }
          }}
        />
      </div>
    </nav>
  );
};

export default Topbar;
