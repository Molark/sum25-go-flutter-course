import 'package:flutter/material.dart';

class ProfileCard extends StatelessWidget {
  final String name;
  final String email;
  final int age;
  final String? avatarUrl;

  const ProfileCard({
    Key? key,
    required this.name,
    required this.email,
    required this.age,
    this.avatarUrl,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    // TODO: Implement profile card UI
    Widget cAvater;
    if (avatarUrl !=null){
      cAvater = CircleAvatar(
            radius: 70,
            backgroundImage: NetworkImage(avatarUrl!),
          );
    } else{
      cAvater = CircleAvatar(
        radius: 70,
        backgroundColor: Colors.green,
        child: Text(name[0].toUpperCase()),
      );
    }
    return Container(
      child: Column(
        children: [
          cAvater,
          Text('$name'),
          Text('Age: $age'),
          Text('$email')
        ],
      ),
    );
  }
}
