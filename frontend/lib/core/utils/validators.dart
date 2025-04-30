class Validators {
  static String? validateDescription(String? value) {
    if (value == null || value.trim().isEmpty) {
      return 'Descrição obrigatória';
    }
    final regex = RegExp(r'^[a-zA-Z0-9 ]+$');
    if (!regex.hasMatch(value.trim())) {
      return 'Somente letras e números';
    }
    if (value.trim().length > 50) {
      return 'Máximo de 50 caracteres';
    }
    return null;
  }

  static String? validateValue(String? value) {
    if (value == null || value.trim().isEmpty) {
      return 'Informe um valor';
    }

    final cleanValue = value.replaceAll(',', '').replaceAll('\$', '').trim();
    final double? number = double.tryParse(cleanValue);

    if (number == null || number <= 0 || number > 99999.99) {
      return 'Valor inválido (máximo 99.999,99)';
    }

    return null;
  }

  static String? validateTransactionId(String? value) {
    if (value == null || value.trim().isEmpty) {
      return 'Informe o ID da transação';
    }

    final regex = RegExp(
      r'^[0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}$',
    );
    if (!regex.hasMatch(value.trim())) {
      return 'ID da transação inválido (UUID esperado)';
    }

    return null;
  }

  static String? validateTransactionDate(DateTime? date) {
    if (date == null) {
      return 'Data da transação obrigatória';
    }

    final today = DateTime.now();
    final lastYear = DateTime(today.year - 5, today.month, today.day);

    if (date.isBefore(lastYear)) {
      return 'Data muito antiga (máx. 5 ano)';
    }

    if (date.isAfter(today)) {
      return 'Data futura não permitida';
    }

    return null;
  }
}
