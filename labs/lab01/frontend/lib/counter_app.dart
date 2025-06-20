import 'package:flutter/material.dart';

class CounterApp extends StatefulWidget {
  const CounterApp({Key? key}) : super(key: key);

  @override
  State<CounterApp> createState() => _CounterAppState();
}

class _CounterAppState extends State<CounterApp> {
  int _counter = 0;

  void _increment() {
    // TODO: Implement increment
    setState(() {
      _counter+=1;
    });
    
  }

  void _decrement() {
    // TODO: Implement decrement
    setState(() {
      _counter-=1;
    });
  }

  void _reset() {
    // TODO: Implement reset
    setState(() {
      _counter=0;
    });
  }

  @override
  Widget build(BuildContext context) {
    // TODO: Implement counter UI   
    return Container(
      child: Column(
        children: [
          Text('$_counter'),
          IconButton(
            icon: const Icon(Icons.add),
            onPressed: _increment,
          ),
          IconButton(
            icon: const Icon(Icons.remove),
            onPressed: _decrement,
          ),
          IconButton(
            icon: const Icon(Icons.refresh),
            onPressed: _reset,
          ),
        ],
    ),
    );
      
  }
}
