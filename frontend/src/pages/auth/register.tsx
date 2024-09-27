import { RegisterForm } from "@/components/auth/RegisterForm";
import { Card, CardHeader, CardTitle } from "@/components/ui/card";

export const RegisterPage = () => {
  return (
    <div className="h-[90vh] flex items-center justify-center">
      <Card className="min-w-[400px] min-h-[400px]">
        <CardHeader>
          <CardTitle className="text-3xl">Register Account</CardTitle>
        </CardHeader>
        <RegisterForm />
      </Card>
    </div>
  );
};
