import argparse


def convert_file(file: str):
    with open(file) as f:
        for line in f.readlines():
            line = line.strip()
            if len(line) == 0:
                print()
                continue
            parts = line.split(' - ')
            if len(parts) != 2:
                raise RuntimeError(f'invalid line: {line}')
            term, translations_line = parts[0], parts[1]
            translations = translations_line.split(', ')
            if len(translations[0]) == 0:
                raise RuntimeError(f'translations not found: {line}')
            joined_translations = '", "'.join(translations)
            print(f'"{term}": {{"{joined_translations}"}},')


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Convert raw vocabulary to GoLang collection.')
    parser.add_argument('files', metavar='F', type=str, nargs='+', help='Input files')

    args = parser.parse_args()
    for file in args.files:
        convert_file(file)
