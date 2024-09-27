import { LoginForm } from "@/components/auth/LoginForm";
import { Card, CardHeader, CardTitle } from "@/components/ui/card";

export const LoginPage = () => {
  return (
    <div className="h-[90vh] flex items-center justify-center">
      <Card className="min-w-[400px] min-h-[300px]">
        <CardHeader>
          <CardTitle className="text-3xl">Login</CardTitle>
        </CardHeader>
        <LoginForm />
      </Card>
    </div>
  );
};
