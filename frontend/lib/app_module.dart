import 'package:flutter_modular/flutter_modular.dart';
import 'package:frontend/modules/home/home_module.dart';
import 'package:frontend/modules/transaction/transaction_module.dart';

class AppModule extends Module {
  @override
  void binds(Injector i) {}

  @override
  void routes(RouteManager r) {
    r.module('/', module: HomeModule());
    r.module('/transaction', module: TransactionModule());
  }
}
