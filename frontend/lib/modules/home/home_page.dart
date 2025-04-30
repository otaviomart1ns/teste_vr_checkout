import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:frontend/core/widgets/button.dart';
import 'package:frontend/core/widgets/app_bar.dart';

class HomePage extends StatelessWidget {
  const HomePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const VRAppBar(),
      body: Padding(
        padding: const EdgeInsets.all(24.0),
        child: Center(
          child: ConstrainedBox(
            constraints: const BoxConstraints(
              maxWidth: 400,
            ), // <= Limite horizontal
            child: Column(
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                const SizedBox(height: 160),
                VRButton(
                  icon: Icons.add,
                  label: 'Nova Transação',
                  onTap: () => Modular.to.pushNamed('/transaction/create'),
                  type: VRButtonType.primary,
                ),
                const SizedBox(height: 20),
                VRButton(
                  icon: Icons.search,
                  label: 'Consultar Transação',
                  onTap: () => Modular.to.pushNamed('/transaction/view'),
                  type: VRButtonType.primary,
                ),
                const SizedBox(height: 20),
                VRButton(
                  icon: Icons.list,
                  label: 'Transações Pendentes',
                  onTap: () => Modular.to.pushNamed('/transaction/pending'),
                  type: VRButtonType.primary,
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
