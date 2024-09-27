import { AuthHeader } from "@/components/layout/AuthHeader";
import { Outlet } from "react-router-dom";

export const AuthLayout = () => {
  return (
    <div className="container mx-auto sm:px-0 px-4">
      <AuthHeader />
      <Outlet />
    </div>
  );
};
