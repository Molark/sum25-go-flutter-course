import 'package:flutter/material.dart';

class RegistrationForm extends StatefulWidget {
  const RegistrationForm({Key? key}) : super(key: key);

  @override
  State<RegistrationForm> createState() => _RegistrationFormState();
}

class _RegistrationFormState extends State<RegistrationForm> {
  final _formKey = GlobalKey<FormState>();
  final _nameController = TextEditingController();
  final _emailController = TextEditingController();
  final _passwordController = TextEditingController();
  bool _isSucess = false;

  @override
  void dispose() {
    _nameController.dispose();
    _emailController.dispose();
    _passwordController.dispose();
    super.dispose();
  }

  void _submitForm() {
    // TODO: Implement form submission
    if (_formKey.currentState!.validate()){
      setState(() {
         _isSucess = true;
      });
    }

  }

  @override
  Widget build(BuildContext context) {
    // TODO: Implement registration form UI
    return Form(
      key: _formKey,
      child: Column(
        children: [
          TextFormField(
            key: const Key('name'),
            controller: _nameController,
            validator: (value){
              if (value == null || value.isEmpty){
                return 'Please enter your name';
              }
              return null;
            }
          ),
          TextFormField(
            key: const Key('email'),
            controller: _emailController,
            validator: (value){
              if (value == null || value.isEmpty || !value.contains("@") || !value.contains(".")){
                return 'Please enter a valid email';
              }
              return null;
            },
          ),
          TextFormField(
            key: const Key('password'),
            controller: _passwordController,
            validator: (value){
              if (value == null || value.isEmpty || value.length < 6){
                return 'Password must be at least 6 characters';
              }
              return null;
            },
          ),
           ElevatedButton(
            onPressed: _submitForm,
            child: Text("Submit"),
          ),
          if (_isSucess)
              Text(
                'Registration successful!',
              key: Key('success'),
              ), 
        ],
      ),
    );
  }
}
