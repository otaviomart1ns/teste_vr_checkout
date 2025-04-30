import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:frontend/modules/transaction/shared/models/local_transaction.dart';
import 'package:hive_flutter/hive_flutter.dart';
import 'app_module.dart';
import 'app_widget.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  await dotenv.load();

  await Hive.initFlutter();
  Hive.registerAdapter(LocalTransactionAdapter());

  runApp(ModularApp(module: AppModule(), child: AppWidget()));
}
