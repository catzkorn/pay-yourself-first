import * as React from "react";
import { Link } from "react-router-dom";

function Nav(): JSX.Element {
  return (
    <div className="bg-gray-100 font-sans w-full m-0">
      <div className="bg-white shadow">
        <div className="container mx-auto px-4">
          <div className="flex items-center justify-between py-4">
            <div className="hidden sm:flex sm:items-center">
              <Link
                to="/"
                className="text-gray-800 text-sm font-semibold hover:text-purple-600 mr-4"
              >
                Budget
              </Link>
              <Link
                to="/income"
                className="text-gray-800 text-sm font-semibold hover:text-purple-600 mr-4"
              >
                Income
              </Link>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Nav;
